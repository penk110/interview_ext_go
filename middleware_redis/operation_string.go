package middleware_redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type StringOperation struct {
	redisClient *redis.Client
	ctx         context.Context
}

func NewStringOperation(ctx context.Context, redisClient *redis.Client) *StringOperation {
	return &StringOperation{
		ctx:         ctx,
		redisClient: redisClient,
	}
}

func (so *StringOperation) Set(key string, v interface{}, attrs ...*OperationAttr) *InterfaceResult {

	// 处理attrs
	expr := OperationAttrs(attrs).Find(AttrExpr).UnWarpOr(0).(time.Duration)

	nx := OperationAttrs(attrs).Find(AttrNX).UnWarpOr(nil)
	if nx != nil {
		return NewInterfaceResult(so.redisClient.SetNX(so.ctx, key, v, expr).Result())
	}

	xx := OperationAttrs(attrs).Find(AttrNX).UnWarpOr(nil)
	if xx != nil {
		return NewInterfaceResult(so.redisClient.SetXX(so.ctx, key, v, expr).Result())
	}

	return NewInterfaceResult(so.redisClient.Set(so.ctx, key, v, expr).Result())
}

func (so *StringOperation) Get(key string) *StringResult {

	return NewStringResult(so.redisClient.Get(so.ctx, key).Result())
}

func (so *StringOperation) MGet(keys ...string) *SliceResult {

	return NewSliceResult(so.redisClient.MGet(so.ctx, keys...).Result())
}
