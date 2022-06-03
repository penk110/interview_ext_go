package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/penk110/interview_ext_go/auth_casbin/pkg/http"
)

var EmployeeH = &employeeH{}

type employeeH struct {
}

// Get get employee detail
func (employee *employeeH) Get(ctx *gin.Context) {

	http.JsonResp(ctx, "ok", "success", "员工详情")
	return
}

// List get employee list
func (employee *employeeH) List(ctx *gin.Context) {

	http.JsonResp(ctx, "ok", "success", "员工列表")
	return
}

// Update employee
func (employee *employeeH) Update(ctx *gin.Context) {

	http.JsonResp(ctx, "ok", "success", "更新员工信息")
	return
}

// Create employee
func (employee *employeeH) Create(ctx *gin.Context) {

	http.JsonResp(ctx, "ok", "success", "添加员工")
	return
}
