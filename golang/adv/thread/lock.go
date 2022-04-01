package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 本代码等效于chn.go

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println("counter =", counter)
	lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go Count(lock)
	}
	for {
		lock.Lock()
		c := counter
		lock.Unlock()

		runtime.Gosched() // 出让时间片

		if c >= 10 {
			break
		}
	}
}
