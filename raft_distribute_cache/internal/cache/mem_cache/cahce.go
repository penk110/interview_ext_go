package mem_cache

import (
	"sync"
	"time"

	"github.com/penk110/interview_ext_go/raft_distribute_cache/internal/cache"
)

type MemCache struct {
	m       map[string]*KV
	keys    []string
	ttlKeys []string
	lock    *sync.RWMutex
	count   int64
}

// KV key-value item TODO: 字段补充; 异步检查是否超时和删除超时key
type KV struct {
	k   string
	v   interface{}
	ttl time.Duration
}

func NewMemCache() cache.Cache {
	memCache := &MemCache{
		m:       map[string]*KV{},
		keys:    []string{},
		ttlKeys: []string{},
		lock:    &sync.RWMutex{},
		count:   0,
	}
	return memCache
}

func (m *MemCache) Keys(size int32) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemCache) KeysWithPrefix(prefix string, size int) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemCache) SetItem(key string, value string) error {
	//TODO implement me
	panic("implement me")
}

func (m *MemCache) SetItemWithTTl(key string, value string, ttl time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (m *MemCache) GetItem(key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MemCache) DelItem(key string) error {
	//TODO implement me
	panic("implement me")
}

func (m *MemCache) Clean() error {
	//TODO implement me
	panic("implement me")
}

func (m *MemCache) cleanExpireKeys() {
}
