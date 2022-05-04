package middleware_redis

import (
	"github.com/go-redis/redis/v8"
	"reflect"
	"testing"
)

func TestRedis(t *testing.T) {
	tests := []struct {
		name string
		want *redis.Client
	}{
		{
			name: "redis client",
			want: Redis(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Redis(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Redis() = %v, want %v", got, tt.want)
			}
		})
	}
}
