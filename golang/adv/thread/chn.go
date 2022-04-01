package main

import (
	"fmt"
	"runtime"
)

// 本代码等效于lock.go

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counted:", ch) // 此句可能执行不到，如果没有runtime.Gosched()
}

func main() {
	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	var count = 0
	for _, ch := range chs {
		v := <-ch
		count += v
		fmt.Println("count =", count)
	}
	runtime.Gosched()
}
