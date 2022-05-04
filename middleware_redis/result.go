package middleware_redis

type Result interface {
	UnWarp() interface{}
}
