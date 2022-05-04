package middleware_redis

type Iterator struct {
	data  []interface{}
	index int
}

func NewIterator(data []interface{}) *Iterator {
	return &Iterator{data: data}
}
func (iterator *Iterator) HasNext() bool {
	if iterator.data == nil || len(iterator.data) == 0 {
		return false
	}
	return iterator.index < len(iterator.data)
}
func (iterator *Iterator) Next() (ret interface{}) {
	ret = iterator.data[iterator.index]
	iterator.index = iterator.index + 1

	return
}
