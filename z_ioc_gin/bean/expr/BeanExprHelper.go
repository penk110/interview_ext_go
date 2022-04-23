package expr

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/penk110/interview_ext_go/z_ioc_gin/bean/expr_lib"
)

func getValueByTokenType(t int, text string, this *beanExprListener) reflect.Value {
	var value reflect.Value
	switch t {
	case expr_lib.BeanExprLexerStringArg:
		stringArg := strings.Trim(text, "'")
		value = reflect.ValueOf(stringArg)
		break
	case expr_lib.BeanExprLexerIntArg:
		v, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			panic("parse int64 error")
		}
		value = reflect.ValueOf(v)
		break
	case expr_lib.BeanExprLexerFloatArg:
		v, err := strconv.ParseFloat(text, 64)
		if err != nil {
			panic("parse float64 error")
		}
		value = reflect.ValueOf(v)
		break
	case expr_lib.BeanExprLexerBoolArg:
		if text == "true" {
			value = reflect.ValueOf(true)
		} else {
			value = reflect.ValueOf(false)
		}
		break
	case expr_lib.BeanExprLexerNilArg:
		value = reflect.ValueOf(nil)
		break
	case expr_lib.BeanExprLexerFuncName | expr_lib.BeanExprLexerMethodName: // 方法
		mRet := BeanExpr(text, this.exprMap)
		if mRet.IsEmpty() {
			value = reflect.ValueOf(nil)
		} else {
			value = reflect.ValueOf(mRet[0])
		}
		break
	default:
		return reflect.Value{}
	}
	return value
}
