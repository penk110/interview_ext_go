package bean

import (
	"log"
	"reflect"

	"github.com/penk110/interview_ext_go/z_ioc_gin/bean/expr"
)

func init() {
	// Factory 初始化一个
	Factory = NewFactory()
}

// Factory mapper factory
var Factory *factory

func NewFactory() *factory {
	return &factory{
		mapper:  make(Mapper),
		ExprMap: make(map[string]interface{}),
	}
}

type factory struct {
	mapper  Mapper
	ExprMap map[string]interface{}
}

func (f *factory) Set(ms ...interface{}) {
	if ms == nil || len(ms) == 0 {
		return
	}
	for _, m := range ms {
		f.mapper.add(m)
	}
}

func (f *factory) GetMapper() Mapper {
	return f.mapper
}

func (f *factory) Get(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	if m := f.mapper.get(v); m.IsValid() {
		return m.Interface()
	}
	return nil
}

func (f *factory) Config(cfgs ...interface{}) {
	for _, cfg := range cfgs {
		t := reflect.TypeOf(cfg)
		if t.Kind() != reflect.Ptr {
			// panic
			panic("require ptr obj")
		}
		if t.Elem().Kind() != reflect.Struct {
			return
		}
		// cfg 注入自身
		f.Set(cfg)
		//
		f.ExprMap[t.Elem().Name()] = cfg
		log.Printf("[inject] ExprMap: %s", t.Elem().Name())
		f.Apply(cfg)
		v := reflect.ValueOf(cfg)

		// 遍历方法列表
		for i := 0; i < t.NumMethod(); i++ {
			method := v.Method(i)

			//log.Printf("[Config] method num: %d, name: %v", i, method.String())

			// TODO: 带参数？
			callRet := method.Call(nil)

			// len(callRet) == 1
			if callRet != nil && len(callRet) > 0 {
				// log.Printf("[Config] method num: %d, name: %v %v", i, method.String(), callRet)
				f.Set(callRet[0].Interface())
			}
		}
	}
}

// Apply 注入bean
func (f *factory) Apply(bean interface{}) {
	if bean == nil {
		return
	}
	v := reflect.ValueOf(bean)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		log.Printf("struct: %s field: %s, canSet: %v, tag inject: %v\n", v.String(), field.Name, v.Field(i).CanSet(), field.Tag.Get("inject"))

		if v.Field(i).CanSet() && field.Tag.Get("inject") != "" {

			if field.Tag.Get("inject") != "-" { // 多例模式
				log.Printf("[inject] 多例模式 struct: %s field type: %s\n", v.String(), field.Type)

				ret := expr.BeanExpr(field.Tag.Get("inject"), f.ExprMap)

				if ret != nil && !ret.IsEmpty() {
					retValue := ret[0]
					if retValue != nil {
						v.Field(i).Set(reflect.ValueOf(retValue))
						f.Apply(retValue)
					}
				}

			} else { // 单例模式
				// 是否存在实例？
				if m := f.Get(field.Type); m != nil {
					// bean 获取类型与初始化
					log.Printf("[inject] 单例模式 struct: %s field type: %s\n", v.String(), field.Type)
					v.Field(i).Set(reflect.ValueOf(m))
					f.Apply(m)
				}
			}
		}
	}
}
