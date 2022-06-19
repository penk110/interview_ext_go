package cache

import "time"

type Cache interface {
	Keys(size int32) ([]string, error)
	KeysWithPrefix(prefix string, size int) ([]string, error)
	SetItem(key string, value string) error
	SetItemWithTTl(key string, value string, ttl time.Duration) error
	GetItem(key string) (string, error)
	DelItem(key string) error
	Clean() error
}
