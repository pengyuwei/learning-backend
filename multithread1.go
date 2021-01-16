package main

import (
  "fmt"
  "runtime"
)

var quit chan int = make(chan int)

func loop() {
  for i := 0; i < 100; i++ {
      fmt.Printf("%d ", i)
  }
  fmt.Printf("\n")
  quit <- 0
}

func main() {
  runtime.GOMAXPROCS(2) // 1-4，看区别

  go loop()
  go loop()

  for i := 0; i < 2; i++ {
      <-quit
  }
}