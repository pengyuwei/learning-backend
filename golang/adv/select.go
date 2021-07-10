package main

import (
	"fmt"
	"time"
)

func ping1(c chan string) {
    time.Sleep(time.Second * 1)
    c <- "ping on channel 1"
}

func ping2(c chan string) {
    time.Sleep(time.Second * 1)
    c <- "ping on channel 2"
}

func useSelect() {
    chan1 := make(chan string)
    chan2 := make(chan string)
    stop := make(chan bool) // 退出通道

    go ping1(chan1)
    go ping2(chan2)
    go func() {
        time.Sleep(time.Second * 1)
        stop <- true
    }()

    // 如果想持续接收消息，可以将select放在循环中
    select { // 阻塞。收到一个消息后，结束阻塞。第二个消息被丢弃。
    case msg1 := <-chan1:
        fmt.Println("recv ", msg1)
    case msg2 := <- chan2:
        fmt.Println("recv ", msg2)
    case <-stop: // 用于主动退出
        fmt.Println("quited. ")
        return
    case <-time.After(1000 * time.Millisecond): // timeout
        fmt.Println("no message received.")
    }
}

func main() {
    useSelect() // 有时1，有时2，有时timeout，有时主动退出
}