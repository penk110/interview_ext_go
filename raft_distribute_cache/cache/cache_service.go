package cache

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/raft"
	jsonIter "github.com/json-iterator/go"

	"github.com/penk110/interview_ext_go/logging_zap"
	"github.com/penk110/interview_ext_go/raft_distribute_cache/internal/cache"
	"github.com/penk110/interview_ext_go/raft_distribute_cache/internal/raft_node"
)

type cacheService struct {
	logger   logging_zap.LoggerImpl
	cache    cache.Cache
	raftNode *raft.Raft // 不方便mock?
}

func NewCacheService(logger logging_zap.LoggerImpl, cache cache.Cache, raftNode *raft.Raft) *cacheService {
	return &cacheService{logger: logger, cache: cache, raftNode: raftNode}
}

func (cs *cacheService) Get(key string) (string, error) {
	value, err := cs.cache.GetItem(key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (cs *cacheService) Set(cacheReq *raft_node.CacheRequest) error {
	reqBytes, _ := jsonIter.Marshal(cacheReq)
	future := cs.raftNode.Apply(reqBytes, time.Second*10)
	if err := future.Error(); err != nil {
		return err
	}
	return nil
}

func (cs *cacheService) Delete(ctx *gin.Context) {

	return
}
