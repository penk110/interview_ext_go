package func_expr

import (
	"fmt"
	"log"

	"github.com/penk110/interview_ext_go/z_ioc_gin/bean/expr"
)

type TestResult struct{}

func (tr *TestResult) Name() string {
	return "test-result"
}
func funcExpr() {
	exprMap := map[string]interface{}{
		"test1": func(name string, age int64) (string, int) {
			log.Println("this is ", name, " and  age is :", age)
			return "shenyi", 19
		},
		"test2": func(str string) string {
			return "test2_" + str
		},
		"test3": func() *TestResult {
			return &TestResult{}
		},
		"test4": func(b bool) (int, bool) {
			return 1, b
		},
		"test5": func(f string) {
			log.Println(f)
		},
	}
	fmt.Println(expr.BeanExpr("test1('test1',19)", exprMap))
	fmt.Println(expr.BeanExpr("test2('test2')", exprMap))
	fmt.Println(expr.BeanExpr("test3()", exprMap)[0].(*TestResult).Name())
	fmt.Println(expr.BeanExpr("test4(true)", exprMap))
	fmt.Println(expr.BeanExpr("test5(test2('golang'))", exprMap))
}
