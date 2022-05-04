package middleware_redis

type SliceResult struct {
	Result []interface{}
	Err    error
}

func NewSliceResult(result []interface{}, err error) *SliceResult {
	return &SliceResult{
		Result: result,
		Err:    err,
	}
}

func (sr *SliceResult) Unwrap() []interface{} {
	if sr.Err != nil {
		panic(sr.Err)
	}
	return sr.Result
}
func (sr *SliceResult) UnwrapOr(v []interface{}) []interface{} {
	if sr.Err != nil {
		return v
	}
	return sr.Result
}

func (sr *SliceResult) Iter() *Iterator {
	return NewIterator(sr.Result)
}
