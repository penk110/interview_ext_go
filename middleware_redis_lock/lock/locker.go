package lock

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"

	middRedis "github.com/penk110/interview_ext_go/middleware_redis"
)

const (
	IncrLuaScript = `
if redis.call('get', KEYS[1]) == ARGV[1] then
  return redis.call('expire', KEYS[1],ARGV[2]) 				
 else
   return '0' 					
end`
)

type Locker struct {
	lockID     string
	key        string
	expire     time.Duration
	release    chan struct{}
	luaScript  *redis.Script // lua 脚本
	resetCount int
}

func NewLocker(key string) *Locker {
	locker := &Locker{
		key:     key,
		release: make(chan struct{}, 1),
	}
	return locker
}

func NewLockerWithAttr(lockID string, key string, atts ...LockerAttrFunc) *Locker {
	locker := &Locker{
		lockID:  lockID,
		key:     key,
		release: make(chan struct{}, 1),
	}
	LockerAttrFuncs(atts).Apply(locker)
	return locker
}

// Chain 链式调用
func (locker *Locker) Chain(atts ...LockerAttrFunc) *Locker {
	LockerAttrFuncs(atts).Apply(locker)
	return locker
}

func (locker *Locker) Lock() (*Locker, error) {
	statusCmd := middRedis.Redis().SetNX(context.Background(), locker.key, "1", locker.expire)
	if ok, err := statusCmd.Result(); err != nil || !ok {
		log.Printf("Lock() %s key: %s, ok: %v, err: %v\n", locker.lockID, locker.key, ok, err)
		return nil, err
	}
	// 自动续约
	go locker.exExpire()
	return locker, nil
}

// exExpire 续约
func (locker *Locker) exExpire() {
	ticker := time.NewTicker(time.Duration(locker.expire.Seconds() * 2 / 3))
	for {
		select {
		case <-ticker.C:
			locker.resetCount += 1
			log.Printf("----------exExpire() %s key: %s, resetCount: %d\n", locker.lockID, locker.key, locker.resetCount)
			locker.resetExpire()
		case <-locker.release: // 退出信号，防止goroutine泄露
			log.Printf("----------exExpire() %s release key: %s, resetCount: %d\n", locker.lockID, locker.key, locker.resetCount)
			return
		}
	}
}

// resetExpire 重置过期时间
func (locker *Locker) resetExpire() {
	cmd := locker.luaScript.Run(context.Background(), middRedis.Redis(), []string{locker.key}, 1, locker.expire.Seconds())
	result, err := cmd.Result()
	if err != nil {
		log.Printf("resetExpire() %s key: %s, err: %v\n", locker.lockID, locker.key, err)
		return
	}
	log.Printf("resetExpire() %s key: %s, result: %v\n", locker.lockID, locker.key, result)
}

func (locker *Locker) Unlock() {
	middRedis.Redis().Del(context.Background(), locker.key)
	locker.release <- struct{}{}
}
