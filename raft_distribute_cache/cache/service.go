package cache

import (
	"github.com/gin-gonic/gin"

	"github.com/penk110/interview_ext_go/raft_distribute_cache/internal/raft_node"
)

type ServiceImpl interface {
	Get(key string) (string, error)
	Set(cacheReq *raft_node.CacheRequest) error
	Delete(ctx *gin.Context)
}
