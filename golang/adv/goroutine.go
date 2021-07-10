// Package main shows how to use channel.
package main

import (
	"fmt"
	"time"
)

func slowFunc(c (chan string)) {
    time.Sleep(time.Second * 2)
    fmt.Println("sleep finished.")
    c <- "slowFunc finished"
}

func useOneMsg() {
    c := make(chan string)
    go slowFunc(c)
    
    /* 
    这里需要等待slowFunc执行完毕，可以这样：
    time.Sleep(time.Second * 2)
    但Go的标准玩法是用通道：make(chan type)
    发送消息：chan <- value
    接收消息：var := <- chan
    */
    msg := <-c // 收到消息前阻塞
    fmt.Println(msg)
}

func receiveBuff(c chan string) {
    for msg := range c {
        fmt.Println(msg)
    }
}

func useBuff() {
    messages := make(chan string, 2)
    messages <- "hello"
    messages <- "world"
    close(messages) //关闭通道，禁止再发送。但缓冲的消息会被保留供读取。
    fmt.Println("pushed two messages onto channel with no receivers")
    time.Sleep(time.Second * 1)
    receiveBuff(messages)
}

func pinger(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "ping"
		<-t.C // delay
	}
}

func flowControl() {
    messages := make(chan string)
	go pinger(messages)
	msg := <-messages
	fmt.Println(msg)
}

func flowControl2() {
    messages := make(chan string)
	go pinger(messages)
	for i:=0;i<5;i++ {
		msg := <-messages
		fmt.Println(msg)
	}
}

func channelReader(messages <- chan string) {
    msg := messages
    fmt.Println(msg)
}

func channelWriter(chn chan<- string) {
    chn <- "hello"
}

func channelReaderAndWriter(messages chan string) {
    msg := <- messages
    fmt.Println(msg)
    messages <- "Hello"
}

// main is entance.
func main() {
    useOneMsg()
    useBuff()
    flowControl()
    flowControl2()
}