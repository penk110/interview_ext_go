package middleware_redis

import "log"

type InterfaceResult struct {
	Result interface{}
	Err    error
}

func NewInterfaceResult(result interface{}, err error) *InterfaceResult {
	return &InterfaceResult{Result: result, Err: err}
}

func (ir *InterfaceResult) UnWarp() interface{} {
	if ir.Err != nil {
		panic(ir.Err)
	}
	return ir.Result
}

// UnWarpOr return default
func (ir *InterfaceResult) UnWarpOr(v interface{}) interface{} {
	if ir.Err != nil {
		log.Printf("[UnWarpOr] err: %s\n", ir.Err.Error())
		return v
	}
	return ir.Result
}

func (ir *InterfaceResult) UnWarpOr2(f func() string) interface{} {
	if ir.Err != nil {
		return f()
	}
	return ir.Result
}
