package middleware

import (
	"github.com/gin-gonic/gin"
)

// Middlewares middleware funcs
func Middlewares() []gin.HandlerFunc {
	fs := []gin.HandlerFunc{
		CheckLogin(),
		RBAC(),
	}
	return fs
}
