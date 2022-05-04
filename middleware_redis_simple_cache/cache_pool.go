package middleware_redis_simple_cache

import (
	"context"
	"github.com/penk110/interview_ext_go/middleware_redis"
	"sync"
	"time"
)

var NewCachePool *sync.Pool

func init() {
	NewCachePool = &sync.Pool{
		New: func() interface{} {
			return NewSimpleCache(
				middleware_redis.NewStringOperation(context.TODO(),
					middleware_redis.Redis()), time.Second*10)
		},
	}
}

func NewCache() *SimpleCache {
	return NewCachePool.Get().(*SimpleCache)
}

func ReleaseNewsCache(cache *SimpleCache) {
	NewCachePool.Put(cache)
}
