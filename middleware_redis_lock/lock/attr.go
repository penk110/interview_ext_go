package lock

import (
	"github.com/go-redis/redis/v8"
	"time"
)

type LockerAttrFunc func(locker *Locker)

type LockerAttrFuncs []LockerAttrFunc

func (attrFuncs LockerAttrFuncs) Apply(locker *Locker) {
	for _, f := range attrFuncs {
		f(locker)
	}
}

func WhitTTL(ttl time.Duration) LockerAttrFunc {
	return func(locker *Locker) {
		locker.expire = ttl
	}
}

func WhitLuaScript(luaScript string) LockerAttrFunc {
	return func(locker *Locker) {
		locker.luaScript = redis.NewScript(luaScript)
	}
}
