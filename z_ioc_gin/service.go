package main

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/penk110/interview_ext_go/z_ioc_gin/bean"
	"github.com/penk110/interview_ext_go/z_ioc_gin/config"
	"github.com/penk110/interview_ext_go/z_ioc_gin/service"
)

func main() {
	_config := config.NewConfig()

	// 实现准备实例
	bean.Factory.Config(_config) // 展开方法

	//  Factory.Set()
	{
		// 测试 userServices
		userService := service.NewUser()
		bean.Factory.Apply(userService) // 处理依赖

		fmt.Println("userService.Order.Name() --->", userService.Order.Name())

		userService.GetUserInfo(uuid.New().String())

		fmt.Println("userService.DB.GetDSN() --->", userService.DB.GetDSN())

	}
}
