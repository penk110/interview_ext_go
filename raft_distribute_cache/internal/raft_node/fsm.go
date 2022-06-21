package raft_node

import (
	"io"

	"github.com/hashicorp/raft"
	"github.com/json-iterator/go"

	"github.com/penk110/interview_ext_go/raft_distribute_cache/internal/cache"
)

type FSM struct {
	cache cache.Cache
}

func NewFSM(cache cache.Cache) *FSM {
	return &FSM{cache: cache}
}

// Apply 持久化操作，需要实现异步持久化？
func (fsm *FSM) Apply(log *raft.Log) interface{} {
	cacheReq := NewCacheRequest()

	// json or protobuf
	err := jsoniter.Unmarshal(log.Data, cacheReq)
	if err != nil {
		return err
	}

	return fsm.cache.SetItem(cacheReq.Key, cacheReq.Value)
}

func (fsm *FSM) Snapshot() (snapshot raft.FSMSnapshot, err error) {
	return nil, nil

}
func (fsm *FSM) Restore(reader io.ReadCloser) error {

	return nil
}
