package cache

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/raft"
	jsonIter "github.com/json-iterator/go"

	"github.com/penk110/interview_ext_go/logging_zap"
	"github.com/penk110/interview_ext_go/raft_distribute_cache/internal/cache"
	"github.com/penk110/interview_ext_go/raft_distribute_cache/internal/raft_node"
)

type CacherController struct {
	logger   logging_zap.LoggerImpl
	cache    cache.Cache
	raftNode *raft.Raft // 不方便mock?
}

func NewCacherController(logger logging_zap.LoggerImpl, cache cache.Cache, raftNode *raft.Raft) *CacherController {
	return &CacherController{
		logger:   logger,
		cache:    cache,
		raftNode: raftNode,
	}
}

func (c *CacherController) Get(ctx *gin.Context) {
	key := ctx.Param("key")
	value, err := c.cache.GetItem(key)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   400000,
			"errMsg": err.Error(),
			"data":   nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200000,
		"message": "ok",
		"data":    value,
	})
	return
}

func (c *CacherController) Set(ctx *gin.Context) {
	cacheReq := raft_node.NewCacheRequest()
	err := ctx.BindJSON(cacheReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":   400000,
			"errMsg": err.Error(),
			"data":   nil,
		})
		return
	}
	reqBytes, _ := jsonIter.Marshal(cacheReq)
	future := c.raftNode.Apply(reqBytes, time.Second*10)
	if err := future.Error(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400000,
			"message": err.Error(),
			"data":    nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200000,
			"message": "ok",
			"data":    cacheReq,
		})
	}

	return
}

func (c *CacherController) Delete(ctx *gin.Context) {

	return
}
