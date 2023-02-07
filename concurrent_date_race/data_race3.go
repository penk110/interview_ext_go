package main

import (
	"fmt"
	"time"
)

var count3 int64

func sum3() {
	for {
		count3++
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go sum3()

	for {
		fmt.Println("--- count: ", count3)
		time.Sleep(time.Second)
	}
}
