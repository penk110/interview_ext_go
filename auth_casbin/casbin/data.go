package casbin

import (
	"github.com/penk110/interview_ext_go/auth_casbin/pkg/models"
	"github.com/penk110/interview_ext_go/auth_casbin/pkg/storage"
	"log"
)

type RoleRel struct {
	PRole  string
	Role   string
	Domain string // 租户域
}

func (roleRel *RoleRel) String() string {
	return roleRel.PRole + ":" + roleRel.Role
}

// GetAllTenants get all tenants
func GetAllTenants() ([]*models.Tenant, error) {
	var tenants []*models.Tenant
	err := storage.DB().
		Find(&tenants).
		Error
	if err != nil {
		return nil, err
	}

	return tenants, nil
}

func getRolesWithDomain(pid int, pname string, roleRels *[]*RoleRel, tenant *models.Tenant) error {
	proles := make([]*models.Role, 0)
	// 根据每个租户ID获取
	log.Printf("role_pid=%v and tenant_id=%v\n", pid, tenant.TenantId)
	err := storage.DB().
		Where("role_pid=? and tenant_id=?", pid, tenant.TenantId).
		Find(&proles).
		Error
	if err != nil {
		return err
	}
	if len(proles) == 0 {
		return nil
	}

	for _, prole := range proles {
		if pname != "" {
			roleRel := &RoleRel{pname, prole.RoleName, tenant.TenantName}
			*roleRels = append(*roleRels, roleRel)
		}
		err := getRolesWithDomain(prole.RoleId, prole.RoleName, roleRels, tenant)
		if err != nil {
			return err
		}

	}

	return nil
}

func GetRolesWithDomain() ([]*RoleRel, error) {
	tenants, err := GetAllTenants()
	if err != nil {
		return nil, err
	}
	log.Println("tenants: ", tenants)
	roleRels := make([]*RoleRel, 0)
	for _, tenant := range tenants {
		err := getRolesWithDomain(0, "", &roleRels, tenant)
		if err != nil {
			return nil, err
		}
		roleRels = append(roleRels, roleRels...)
	}
	return roleRels, nil
}

func GetUserRolesWithDomain() ([]*models.Users, error) {
	var users []*models.Users
	err := storage.DB().
		Select(" us.user_name, rs.role_name,t.tenant_name ").
		Table(" users us,user_roles ur ,roles rs, tenants t ").
		Where(" us.user_id=ur.user_id and ur.role_id=rs.role_id and rs.tenant_id=t.tenant_id ").
		Find(&users).
		Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetRouterRolesWithDomain() ([]*models.Routers, error) {
	var routers []*models.Routers
	err := storage.DB().Select(" route.r_uri,route.r_method,r.role_name,t.tenant_name ").
		Table(" routers route,router_roles rs,roles r ,tenants t ").
		Where(" route.r_id=rs.router_id and rs.role_id=r.role_id and r.tenant_id=t.tenant_id ").
		Find(&routers).
		Error
	if err != nil {
		return nil, err
	}
	return routers, nil
}
