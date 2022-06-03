package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/penk110/interview_ext_go/auth_casbin/casbin"
	"github.com/penk110/interview_ext_go/auth_casbin/conf"
	"github.com/penk110/interview_ext_go/auth_casbin/handler"
	"github.com/penk110/interview_ext_go/auth_casbin/middleware"
	"github.com/penk110/interview_ext_go/auth_casbin/pkg/storage"
)

func main() {
	var (
		cfgPath string
		cfg     *conf.Conf

		err error
	)

	// init conf
	if cfg, err = conf.Init(cfgPath); err != nil {
		log.Fatal(err)
	}
	// init storage
	if err = storage.Init(cfg); err != nil {
		log.Fatal(err)
	}
	// init casbin
	if err = casbin.Init(cfg); err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	// use middleware
	r.Use(middleware.Middlewares()...)

	deptGroup := r.Group("/api/v1/dept")
	{
		deptGroup.GET(":dept_id", handler.DeptH.Get)
		deptGroup.GET("", handler.DeptH.List)
		deptGroup.PUT(":dept_id", handler.DeptH.Update)
		deptGroup.POST("", handler.DeptH.Create)
	}
	employeeGroup := r.Group("/api/v1/employee")
	{
		employeeGroup.GET(":employee_id", handler.EmployeeH.Get)
		employeeGroup.GET("", handler.EmployeeH.List)
		employeeGroup.PUT(":employee_id", handler.EmployeeH.Update)
		employeeGroup.POST("", handler.EmployeeH.Create)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
		return
	}
}
