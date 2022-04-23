package expr

import (
	"log"
	"reflect"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/penk110/interview_ext_go/z_ioc_gin/bean/expr_lib"
)

type ResultSet []interface{}

func BeanExpr(expr string, exprMap map[string]interface{}) ResultSet {
	is := antlr.NewInputStream(expr)

	lexer := expr_lib.NewBeanExprLexer(is)
	ts := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := expr_lib.NewBeanExprParser(ts)

	lis := &beanExprListener{exprMap: exprMap}
	antlr.ParseTreeWalkerDefault.Walk(lis, p.Start())

	return lis.Run()
}

func newResultSet() ResultSet {
	return make(ResultSet, 0)
}
func (rs ResultSet) IsEmpty() bool {
	return len(rs) == 0
}
func (rs ResultSet) Len() int {
	return len(rs)
}
func result(values []reflect.Value) ResultSet {
	ret := newResultSet()
	if values == nil || len(values) == 0 {
		return ret
	}
	for _, v := range values {
		ret = append(ret, v.Interface())
	}
	return ret
}

type beanExprListener struct {
	*expr_lib.BaseBeanExprListener
	funcName   string
	args       []reflect.Value
	methodName string // 方法名
	execType   uint8  // 执行类型  0代表函数,默认值 1代表struct执行
	exprMap    map[string]interface{}
}

func (rs *beanExprListener) ExitMethodCall(ctx *expr_lib.MethodCallContext) {
	rs.execType = 1
	rs.methodName = ctx.GetStart().GetText()

}
func (rs *beanExprListener) ExitFuncCall(ctx *expr_lib.FuncCallContext) {
	rs.funcName = ctx.GetStart().GetText()
}
func (rs *beanExprListener) ExitFuncArgs(ctx *expr_lib.FuncArgsContext) {
	for i := 0; i < ctx.GetChildCount(); i++ {
		if token, ok := ctx.GetChild(i).GetPayload().(*antlr.BaseParserRuleContext); ok {

			value := getValueByTokenType(token.GetStart().GetTokenType(), token.GetText(), rs)
			if value.IsValid() {
				rs.args = append(rs.args, value)
			}
		}

		//a:=ctx.GetChild(i).GetPayload().(*antlr.BaseParserRuleContext)
		//
		//value:=getValueByTokenType(token.GetTokenType(),token.GetText())
		//if value.IsValid(){
		//	rs.args = append(rs.args, value)
		//}
		//if a,ok:=ctx.GetChild(i).GetPayload().(*antlr.CommonToken);ok{
		//	log.Println("111",a.GetText())
		//}
		//fmt.Printf("%T\n",ctx.GetChild(i).GetPayload())

	}

}
func (rs *beanExprListener) findField(method string, v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if field := v.FieldByName(method); field.IsValid() {
		return field
	}
	return reflect.Value{}
}
func (rs *beanExprListener) Run() ResultSet {
	if rs.exprMap == nil {
		panic("exprMap required")
	}
	switch rs.execType {
	case 0: // 默认
		if f, ok := rs.exprMap[rs.funcName]; ok {
			v := reflect.ValueOf(f)
			if v.Kind() == reflect.Func {
				return result(v.Call(rs.args))
			}
		}
		break
	case 1: // struct方法执行
		ms := strings.Split(rs.methodName, ".")
		if obj, ok := rs.exprMap[ms[0]]; ok {
			objValue := reflect.ValueOf(obj)
			current := objValue
			for i := 1; i < len(ms); i++ {
				if i == len(ms)-1 { // 最后一个是方法名
					if method := current.MethodByName(ms[i]); !method.IsValid() {
						panic("method error:" + ms[i])
					} else {
						return result(method.Call(rs.args))
					}
				}
				field := rs.findField(ms[i], current)
				if field.IsValid() {
					current = field
				} else {
					panic("field error:" + ms[i])
				}
			}
		}
	default:
		log.Println("nothing to do")
	}
	return newResultSet()

}
