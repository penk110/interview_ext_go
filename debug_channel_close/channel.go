package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var ch chan int
	ch = make(chan int, 100)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			j := i
			ch <- j
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 101; i++ {

			// 往一个已经被取空的通道中取元素得到对应类型的零值
			j := <-ch
			fmt.Printf("func2<%v> read ch: %v\n", i, j)

			// 往已经被关闭的通道中推数据会报错
			//if i > 100 {
			//	ch <- i
			//}
		}
	}()

	fmt.Printf("%v\n", ch)
	wg.Wait()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	defer func() {
		e := recover()
		fmt.Println("----defer-----", e)
	}()
	// panic
	ch <- 100
}
