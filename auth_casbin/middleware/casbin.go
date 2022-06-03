package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/penk110/interview_ext_go/auth_casbin/casbin"
	"log"
	"net/http"
)

func RBAC() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, exist := ctx.Get("user_name")
		domain := ctx.Request.Header.Get("domain")
		if !exist {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    "400400",
				"message": "get user_name from ctx header failed",
				"data":    nil,
			})
			return
		}
		log.Printf("user: %s, domain: %s, RequestURI: %s, Method: %s\n", user, domain, ctx.Request.RequestURI, ctx.Request.Method)
		ok, err := casbin.E.Enforce(user, domain, ctx.Request.RequestURI, ctx.Request.Method)
		if !ok || err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    "400404",
				"message": "forbidden",
				"data":    nil,
			},
			)
			return
		}
		ctx.Next()
	}
}
