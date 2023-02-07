package main

import (
	"fmt"
	"time"
)

var count2 int64

var ch chan int64

func sum2() {
	for {
		count2++
		ch <- count2
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	ch = make(chan int64)
	go sum2()

	for c := range ch {

		fmt.Println("--- count: ", c)
		time.Sleep(time.Second)
	}
}
