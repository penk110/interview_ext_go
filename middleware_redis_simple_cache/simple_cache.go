package middleware_redis_simple_cache

import (
	"time"

	"github.com/penk110/interview_ext_go/middleware_redis"
)

type SimpleCache struct {
	Operation *middleware_redis.StringOperation
	Expr      time.Duration
}

func NewSimpleCache(operation *middleware_redis.StringOperation, expr time.Duration) *SimpleCache {
	return &SimpleCache{
		Operation: operation,
		Expr:      expr,
	}
}

func (cache *SimpleCache) SetCache(key string, value interface{}) {

	cache.Operation.Set(key, value, middleware_redis.WithExpr(cache.Expr))
}

func (cache *SimpleCache) GetCache(key string, getFunc func() string) (ret interface{}) {
	ret = cache.Operation.Get(key).UnWarpOr2(getFunc)
	cache.SetCache(key, ret)
	return
}
