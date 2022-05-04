package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/penk110/interview_ext_go/middleware_redis_simple_cache"
)

func main() {
	engine := gin.Default()
	{
		engine.GET("/api/v1/step_records/:task_id", GetStepRecord)
	}

	err := engine.Run(":8080")
	if err != nil {
		panic(err)
		return
	}
}

func GetStepRecord(ctx *gin.Context) {

	// 1edd80cb606049587ff43e0522df8cf7
	taskId := ctx.Param("task_id")
	log.Println("[GetStepRecord] task_id: " + taskId)
	if taskId == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "400000",
			"msg":  fmt.Errorf("task_id not allow empty"),
			"data": nil,
		})
		return
	}
	// 获取方法的key
	key := "/tmp/cobra/step_records/" + taskId
	var getFunc = func() string {
		result := middleware_redis_simple_cache.GetDB().Database("cobra").
			Collection("step_records").
			FindOne(context.TODO(), map[string]interface{}{"taskid": taskId})
		m := make(map[string]interface{})
		result.Decode(&m)
		b, _ := json.Marshal(m)
		return string(b)
	}

	cacheRet := middleware_redis_simple_cache.NewCache().GetCache(key, getFunc)
	if cacheRet == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": "400000",
			"msg":  "err",
			"data": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": "200000",
		"msg":  "success",
		"data": cacheRet,
	})
	return
}
