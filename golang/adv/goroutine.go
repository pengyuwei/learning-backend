package main

import (
	"fmt"
	"time"
)

func slowFunc() {
    time.Sleep(time.Second * 2)
    fmt.Println("sleep finished.")
}

func main() {
    go slowFunc()
    fmt.Println("Ha Ha")
    /* 
    这里需要等待slowFunc执行完毕，可以这样：
    time.Sleep(time.Second * 2)
    但Go的标准玩法是用通道：make(chan type)
    发送消息：chan <- value
    接收消息：var := <- chan
    */
    c := make(chan string)
    c <- "Hello"
}