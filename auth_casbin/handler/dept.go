package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/penk110/interview_ext_go/auth_casbin/pkg/http"
)

var DeptH = &deptH{}

type deptH struct {
}

// Get get dept detail
func (dept *deptH) Get(ctx *gin.Context) {

	http.JsonResp(ctx, "ok", "success", "部门详情")
	return
}

// List get dept list
func (dept *deptH) List(ctx *gin.Context) {

	http.JsonResp(ctx, "ok", "success", "部门列表")
	return
}

// Update dept
func (dept *deptH) Update(ctx *gin.Context) {

	http.JsonResp(ctx, "ok", "success", "更新部门信息")
	return
}

// Create dept
func (dept *deptH) Create(ctx *gin.Context) {

	http.JsonResp(ctx, "ok", "success", "添加部门")
	return
}
