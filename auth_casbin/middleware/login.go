package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckLogin check login func
func CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 默认先放到header里
		if ctx.Request.Header.Get("token") == "" {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "400400",
				"message": "token empty",
				"data":    nil,
			})
			return
		} else {
			// 解析token，将user put 到ctx中
			ctx.Set("user_name", ctx.Request.Header.Get("token"))
			ctx.Next()
		}
	}
}
