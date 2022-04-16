package case_js_promise

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"
)

type Resolve func(out *Out)

type Reject func(err error)

type PromiseFunc func(resolve Resolve, reject Reject)

type PromiseOpt func(promise *Promise)

type PromiseOpts []PromiseOpt

func (opts PromiseOpts) apply(p *Promise) {
	for _, opt := range opts {
		opt(p)
	}
}

// ----------
// TODO: 预定义opts

func WitchTimeoutOpt(t time.Duration) PromiseOpt {
	return func(promise *Promise) {
		promise.timeout = t
	}
}

// ----------

type Out struct {
	Resp *http.Response
	Err  error
	Cost time.Duration
}

// Promise 对象
type Promise struct {
	resolve Resolve
	reject  Reject
	pf      PromiseFunc // 执行
	retry   int         // default 1
	wg      sync.WaitGroup
	timeout time.Duration
}

func NewPromise(pf PromiseFunc, retry int) *Promise {
	promise := &Promise{
		resolve: nil,
		reject:  nil,
		pf:      pf,
		retry:   0,
		wg:      sync.WaitGroup{},
	}
	if retry > 0 {
		promise.retry = retry
	}
	return promise
}

// WithOpts 添加opts
func (p *Promise) WithOpts(opts ...PromiseOpt) *Promise {
	PromiseOpts(opts).apply(p)
	return p
}

func (p *Promise) Then(resolve Resolve) *Promise {
	p.resolve = resolve
	return p
}

func (p *Promise) Catch(reject Reject) *Promise {
	p.reject = reject
	return p
}

func (p *Promise) Done() {
	doneCh := make(chan struct{})

	timeoutCtx, cancel := context.WithTimeout(context.TODO(), p.timeout)
	defer cancel()

	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		p.pf(p.resolve, p.reject)
	}()
	go func() {
		// done
		defer func() { doneCh <- struct{}{} }()
		p.wg.Wait()
	}()

	select {
	case <-timeoutCtx.Done():
		// TODO: retry ???
		if p.retry > 0 {
			p.retry--
			log.Printf("[Done] timeout retry, has: %d", p.retry)
			p.Done()
		}
		return
	case <-doneCh:
		return
	}

}
