package models

import "fmt"

type Tenant struct {
	TenantId   int    `gorm:"column:tenant_id;primaryKey"`
	TenantName string `gorm:"column:tenant_name"`
}

func (tenant *Tenant) TableName() string {
	return "tenants"
}

func (tenant *Tenant) String() string {
	return fmt.Sprintf("%d:%s", tenant.TenantId, tenant.TenantName)
}
