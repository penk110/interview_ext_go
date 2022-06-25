package cache

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/penk110/interview_ext_go/raft_distribute_cache/internal/raft_node"
)

type CacherController struct {
	cacheService ServiceImpl
}

func NewCacherController(cacheService ServiceImpl) *CacherController {
	return &CacherController{
		cacheService: cacheService,
	}
}

func (c *CacherController) Get(ctx *gin.Context) {
	key := ctx.Param("key")
	value, err := c.cacheService.Get(key)
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
	if err := c.cacheService.Set(cacheReq); err != nil {
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
