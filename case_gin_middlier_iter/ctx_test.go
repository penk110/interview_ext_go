package ctx

import (
	"fmt"
	"testing"
)

/*
	洋葱模型
*/

func middleware1() Handler {
	return func(ctx *Ctx) {
		fmt.Println("middleware 1 start")
		ctx.Next()
		fmt.Println("middleware 1 done")
	}
}

func middleware2() Handler {
	return func(ctx *Ctx) {
		fmt.Println("middleware 2 start")
		//被取消执行了
		//ctx.Abort()

		ctx.Next() // 剥洋葱
		fmt.Println("middleware 2 done")
	}
}

func middleware3() Handler {
	return func(ctx *Ctx) {
		fmt.Println("middleware 3 start")
		ctx.Next()
		fmt.Println("middleware 3 done")
	}
}

func TestCtx(t *testing.T) {
	var ctx = &Ctx{}
	ctx.Use(middleware1())
	ctx.Use(middleware2())
	ctx.Use(middleware3())
	ctx.Get("/index", func(ctx *Ctx) {
		fmt.Println("[GET] index")
		return
	})
	ctx.Run()
}
