package ctx

import (
	"math"
)

const abortIndex int8 = math.MaxInt8 / 2

type ICtx interface {
	Next()
	Abort()
	Use(h Handler)
	Get(path string, h Handler)
	Put(path string, h Handler)
}

type Handler func(ctx *Ctx)

type Ctx struct {
	handlers []Handler
	index    int8
}

func (c *Ctx) Use(h Handler) {
	if c.handlers == nil || len(c.handlers) == 0 {
		c.handlers = []Handler{}
	}
	c.handlers = append(c.handlers, h)
}

func (c *Ctx) Run() {
	c.handlers[0](c) // 执行第一个handler，如果是中间件则先执行中间件 -> next进入下个handler(如果是中途被abort则退出)
}

func (c *Ctx) Next() {
	// 是否被设置退出：c.index = _ABORTIndex
	if c.index < int8(len(c.handlers)) {
		c.index++
		c.handlers[c.index](c)
	}
}

func (c *Ctx) Abort() {
	c.index = abortIndex
}

func (c *Ctx) Get(path string, h Handler) {
	c.handlers = append(c.handlers, h)
	return
}

func (c *Ctx) Put(path string, h Handler) {
	c.handlers = append(c.handlers, h)
	return
}
