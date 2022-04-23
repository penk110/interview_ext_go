package examples

import (
	"fmt"

	"github.com/penk110/interview_ext_go/z_ioc_gin/bean/expr"
)

type UserRole struct {
	RoleName string
}

func (userRole *UserRole) GetRole(prefix string) string {
	return prefix + ":" + userRole.RoleName
}

type User struct {
	Name string
	Role *UserRole
}

func (user *User) GetName() string {
	return user.Name
}

// NewUser new user instance
func NewUser(name string, role string) *User {
	return &User{Name: name, Role: &UserRole{RoleName: role}}
}

func structExpr() {
	// expr map
	exprMap := map[string]interface{}{
		"user": NewUser("gopher", "admin"),
	}
	fmt.Println(expr.BeanExpr("user.GetName()", exprMap))             //方法名 大小写敏感
	fmt.Println(expr.BeanExpr("user.Role.GetRole('当前角色是')", exprMap)) //方法名 大小写敏感
}
