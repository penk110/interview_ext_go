package middleware_redis

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

var once sync.Once

const (
	defaultAddr = "127.0.0.1:6379"
)

var addr string

func init() {
	addr = os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = defaultAddr
	}
}

func Redis() *redis.Client {
	once.Do(func() {
		if client != nil {
			return
		}
		client = redis.NewClient(&redis.Options{
			Network: "tcp",
			Addr:    addr,

			MaxRetries:      0,
			MinRetryBackoff: 8 * time.Millisecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
			MaxRetryBackoff: 512 * time.Millisecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔

			// 超时
			DialTimeout:  30 * time.Second,
			ReadTimeout:  3 * time.Second,
			WriteTimeout: 3 * time.Second,
			PoolTimeout:  3 * time.Second,

			PoolFIFO:           false,
			PoolSize:           50,
			MinIdleConns:       20,
			MaxConnAge:         0,
			IdleTimeout:        5 * time.Minute, //闲置超时
			IdleCheckFrequency: 1 * time.Second,
		})
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
		defer cancel()
		pong, err := client.Ping(ctx).Result()
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Printf("%v", pong)
	})
	return client
}
