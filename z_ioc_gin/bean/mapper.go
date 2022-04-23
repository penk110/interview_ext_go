package bean

import "reflect"

type Mapper map[reflect.Type]reflect.Value

// add 添加bean对象
func (mapper Mapper) add(bean interface{}) {
	t := reflect.TypeOf(bean)
	if t.Kind() != reflect.Ptr {
		panic("require ptr obj")
	}
	mapper[t] = reflect.ValueOf(bean)
}

// get 获取bean对象
func (mapper Mapper) get(bean interface{}) reflect.Value {
	var (
		t reflect.Type
	)
	if bt, ok := bean.(reflect.Type); ok {
		t = bt
	} else {
		t = reflect.TypeOf(bean)
	}
	if v, ok := mapper[t]; ok {
		return v
	}

	// 接口继承
	for k, v := range mapper {
		// 是否继承
		if t.Kind() == reflect.Interface && k.Implements(t) {
			return v
		}
	}
	return reflect.Value{}
}
