package main
import "fmt"


// deadlock example
// fatal error: all goroutines are asleep - deadlock!
 func main() {
     ch := make(chan int)
     ch <- 1 // 1流入信道，堵塞当前线, 没人取走数据信道不会打开
     fmt.Println("This line code wont run") //在此行执行之前Go就会报死锁
 }