package cache

import (
	"github.com/gin-gonic/gin"
)

func (c *CacherController) Register(engine *gin.Engine) {
	engine.POST("/api/v1/cache", c.Set)

	engine.GET("/api/v1/cache/:key", c.Get)
	engine.DELETE("/api/v1/cache/:key", c.Delete)
}
