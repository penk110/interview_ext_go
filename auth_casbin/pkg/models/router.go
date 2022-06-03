package models

import "fmt"

type Routers struct {
	RouterId     int    `gorm:"column:user_id;primaryKey"`
	RouterName   string `gorm:"column:r_name"`
	RouterUri    string `gorm:"column:r_uri;"`
	RouterMethod string `gorm:"column:r_method"`
	RoleName     string
	Domain       string `gorm:"column:tenant_name"`
}

func (routers *Routers) TableName() string {
	return "routers"
}
func (routers *Routers) String() string {
	return fmt.Sprintf("%s-%s", routers.RouterName, routers.RoleName)
}
