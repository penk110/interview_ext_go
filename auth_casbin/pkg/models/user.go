package models

import "fmt"

type Users struct {
	UserID   int    `gorm:"column:user_id;primaryKey"`
	UserName string `gorm:"column:user_name"`
	RoleName string `gorm:"column:role_name"`
	Domain   string `gorm:"column:tenant_name"`
}

func (users *Users) TableName() string {
	return "users"
}
func (users *Users) String() string {
	return fmt.Sprintf("%s-%s-%s", users.UserName, users.RoleName, users.Domain)
}
