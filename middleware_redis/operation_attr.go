package middleware_redis

import (
	"errors"
	"time"
)

const (
	AttrExpr = "expr"
	AttrNX   = "nx"
	AttrXX   = "xx"
)

// args传入

type Empty struct{}

type OperationAttr struct {
	Name  string
	Value interface{}
}

type OperationAttrs []*OperationAttr

func (attrs OperationAttrs) Find(name string) *InterfaceResult {
	for _, attr := range attrs {
		if attr.Name == name {
			return NewInterfaceResult(attr.Value, nil)
		}
	}
	return NewInterfaceResult(nil, errors.New("operation attr not found, name: "+name))
}

func WithExpr(t time.Duration) *OperationAttr {
	return &OperationAttr{
		Name:  AttrExpr,
		Value: t,
	}
}

func WithNX() *OperationAttr {
	return &OperationAttr{
		Name:  AttrNX,
		Value: Empty{},
	}
}

func WithXX() *OperationAttr {
	return &OperationAttr{
		Name:  AttrXX,
		Value: Empty{},
	}
}
