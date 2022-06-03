package casbin

import (
	"fmt"
	"log"
	"path"

	"github.com/casbin/casbin/v2"

	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"github.com/penk110/interview_ext_go/auth_casbin/conf"
	"github.com/penk110/interview_ext_go/auth_casbin/pkg/storage"
)

var E *casbin.Enforcer

func Init(cfg *conf.Conf) error {
	adapterDB := storage.DB()
	adapter, err := gormAdapter.NewAdapterByDB(adapterDB)
	if err != nil {
		return err
	}

	modelFile := path.Join(cfg.Casbin.ResourcePath, "model_t.conf")
	E, err = casbin.NewEnforcer(modelFile, adapter)
	if err != nil {
		fmt.Println(modelFile, err)
		return err
	}

	// 自定义 matcher
	E.AddFunction("methodMatch", func(arguments ...interface{}) (i interface{}, e error) {
		if len(arguments) == 2 {
			k1, k2 := arguments[0].(string), arguments[1].(string)
			return MethodMatch(k1, k2), nil
		}
		return nil, fmt.Errorf("methodMatch error")
	})

	if err = E.LoadPolicy(); err != nil {
		return err
	}

	if err := initPolicyWithDomain(); err != nil {
		return err
	}
	return nil
}

func initPolicyWithDomain() error {
	// 初始化 角色关系

	// 拼接策略
	roles, err := GetRolesWithDomain()
	log.Printf("GetRolesWithDomain() roles: %#v, err: %v\n", roles, err)
	if err != nil {
		return fmt.Errorf("cant find any roles, err: %v", err)
	}

	for _, role := range roles {
		fmt.Println("role ------ ", role.PRole, role.Role, role.Domain)
		_, err := E.AddRoleForUserInDomain(role.PRole, role.Role, role.Domain)
		if err != nil {
			return err
		}
	}

	userRoles, err := GetUserRolesWithDomain()
	log.Printf("GetUserRolesWithDomain() userRoles: %#v, err: %v\n", userRoles, err)
	if err != nil {
		return fmt.Errorf("cant find any user roles, err: %v", err)
	}
	for _, userRole := range userRoles {
		// 指定 domain参数
		fmt.Println("userRole ------  ", userRole.UserName, userRole.RoleName, userRole.Domain)
		_, err := E.AddRoleForUserInDomain(userRole.UserName, userRole.RoleName, userRole.Domain)
		if err != nil {
			return err
		}
	}

	routerRoles, err := GetRouterRolesWithDomain()
	log.Printf("GetRouterRolesWithDomain() routerRoles: %#v, err: %v\n", routerRoles, err)
	if err != nil {
		return fmt.Errorf("cant find any router roles, err: %v", err)
	}
	for _, routerRole := range routerRoles {
		fmt.Println("userRole ------  ", routerRole.RoleName, routerRole.Domain, routerRole.RouterUri, routerRole.RouterMethod)
		_, err := E.AddPolicy(routerRole.RoleName, routerRole.Domain, routerRole.RouterUri, routerRole.RouterMethod)
		if err != nil {
			return err
		}
	}

	return nil
}
