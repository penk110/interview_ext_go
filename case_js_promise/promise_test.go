package case_js_promise

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

const (
	GoCnUrl = "https://gocn.vip/"
)

func TestPromise(t *testing.T) {
	// success
	resolve := Resolve(func(out *Out) {
		// TODO: do something
		start := time.Now()
		out.Resp, out.Err = http.Get(GoCnUrl)
		out.Cost = time.Now().Sub(start)
		log.Printf("[resolve] cost: %d", out.Cost.Milliseconds())
	})

	// failed
	reject := Reject(func(err error) {
		log.Printf("[reject] err: %v", err)
	})

	var out = &Out{}
	pf := func(resolve Resolve, reject Reject) {
		// resolve		reject
		resolve(out)

		//if out.Cost.Seconds() > 1 {
		//	out.Err = fmt.Errorf("cost gather 1s: %f", out.Cost.Seconds())
		//}
		if out.Cost.Milliseconds() > 120 {
			out.Err = fmt.Errorf("cost gather 100ms: %f", out.Cost.Seconds())
		}
		if out.Err != nil {
			reject(out.Err)
		}
	}
	retry := 2

	//NewPromise(pf, retry).Then(resolve).Catch(reject).Done()

	// 添加超时 WithTimeOpt
	NewPromise(pf, retry).WithOpts(WitchTimeoutOpt(time.Millisecond * 120)).Then(resolve).Catch(reject).Done()

	var body []byte
	body, err := ioutil.ReadAll(out.Resp.Body)
	if err != nil {
		t.Errorf("read body failed, err: %d", err)
		return
	}

	t.Logf("body len: %d, data: %s", len(body), body)
}
