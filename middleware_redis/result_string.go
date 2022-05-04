package middleware_redis

import "log"

type StringResult struct {
	Result string
	Err    error
}

func NewStringResult(result string, err error) *StringResult {
	return &StringResult{
		Result: result,
		Err:    err,
	}
}

func (sr *StringResult) UnWarp() interface{} {
	if sr.Err != nil {
		panic(sr.Err)
	}
	return sr.Result
}

// UnWarpOr return default
func (sr *StringResult) UnWarpOr(s string) string {
	if sr.Err != nil {
		return s
	}
	return sr.Result
}

func (sr *StringResult) UnWarpOr2(f func() string) string {
	if sr.Err != nil {
		log.Printf("[UnWarpOr2] use getFunc, err: %v", sr.Err)
		return f()
	}
	return sr.Result
}
