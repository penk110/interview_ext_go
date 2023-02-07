package main

import (
	"fmt"
	"time"
)

var count int64

func sum() {
	for {
		count++
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	go sum()

	for {
		count++
		fmt.Println("--- count: ", count)
		time.Sleep(time.Second)
	}
}
