package main
import "fmt"
import "time"


var complete chan int = make(chan int)

func loop() {
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", i)
    }
}

func loop2() {
    for i := 0; i < 10; i++ {
        fmt.Printf("%d ", i)
    }

    complete <- 0
}


func main_1() {
    loop()
    loop()
}

func main_2() {
    go loop()
    loop()
    time.Sleep(time.Second)
}

func main_3() {
    var messages chan string = make(chan string)
    go func(message string) {
        messages <- message // 存消息
    }("Ping!")

    fmt.Println(<-messages) // 取消息
}

func main_4() {
    // fatal error: all goroutines are asleep - deadlock!
    go loop2()
    loop()
    <- complete
}

func main() {
    main_4()
}