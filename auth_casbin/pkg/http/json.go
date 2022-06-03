package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JsonResp(ctx *gin.Context, code string, msg interface{}, data interface{}) {
	switch code {
	case "ok":
		ctx.JSON(http.StatusOK, gin.H{
			"code":    "200000",
			"message": msg,
			"data":    data,
		},
		)
	case "forbidden":
		ctx.JSON(http.StatusForbidden, gin.H{
			"code":    "400403",
			"message": msg,
			"data":    nil,
		},
		)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    "40000",
			"message": msg,
			"data":    data,
		},
		)
	}
}
