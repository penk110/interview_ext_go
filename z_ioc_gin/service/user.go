package service

import "fmt"

type User struct {
	Order *Order `inject:"-"`
	DB    *DB    `inject:"-"`
}

func NewUser() *User {
	return &User{}
}
func (user *User) GetUserInfo(uid string) {

	fmt.Println("[User] GetUserInfo")
}
func (user *User) GetOrderInfo(uid string) {

	fmt.Println("获取用户ID=", uid, "的订单信息")
}
