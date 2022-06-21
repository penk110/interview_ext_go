package badger_cache

import (
	"time"

	"github.com/dgraph-io/badger/v2"

	"github.com/penk110/interview_ext_go/raft_distribute_cache/internal/cache"
)

type BadgerCache struct {
	*badger.DB
}

func NewBadgerCache(path string) cache.Cache {
	options := badger.DefaultOptions(path)

	// just windows
	options.Truncate = true

	db, err := badger.Open(options)
	if err != nil {
		panic(err)
	}

	badgerCache := &BadgerCache{
		DB: db,
	}

	return badgerCache
}

func (b *BadgerCache) Keys(size int32) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BadgerCache) KeysWithPrefix(prefix string, size int) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BadgerCache) SetItem(key string, value string) error {
	err := b.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
	return err
}

func (b *BadgerCache) SetItemWithTTl(key string, value string, ttl time.Duration) error {
	err := b.Update(func(txn *badger.Txn) error {
		entry := badger.NewEntry([]byte(key), []byte(value)).WithTTL(ttl)
		return txn.SetEntry(entry)
	})
	return err
}

func (b *BadgerCache) GetItem(key string) (string, error) {
	var ret string
	err := b.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			ret = string(val)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return "", err
	}
	return ret, nil
}

func (b *BadgerCache) DelItem(key string) error {
	//TODO implement me
	panic("implement me")
}

func (b *BadgerCache) Clean() error {
	//TODO implement me
	panic("implement me")
}
