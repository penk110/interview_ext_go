package service

import "fmt"

type IOrder interface {
	Name() string
}

type Order struct {
	Version string
	DB      *DB `inject:"-"`
}

func NewOrder() *Order {
	return &Order{
		Version: "1.0",
	}
}
func (order *Order) GetOrderInfo(uid int) {

	fmt.Println("获取用户ID=", uid, "的订单信息")
}
func (order *Order) Name() string {

	return "order"
}
