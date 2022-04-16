package channel_recive_result

import (
	"testing"
	"time"
)

func TestLooper(t *testing.T) {
	looper := InitDefaultLooper()
	for i := 0; i < 10; i++ {
		looper.Push(NewJob())
	}
	time.Sleep(time.Second * 5)
}
