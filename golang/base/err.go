package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

// 不支持传统的try-catch-finally结构
// Go的习惯：将错误返回给上层处理

func useErr() {
    // var file []byte
    // var err error
    file, err := ioutil.ReadFile("nofile.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%s\n", file)
}

func newError() {
    err := errors.New("A new error")
    if err != nil {
        fmt.Println(err)
    }
}

func fmtError() {
    name, role := "Tom", "Mary"
    err := fmt.Errorf("The %v %v", role, name)
    if err != nil {
        fmt.Println(err)
    }
}

func retError(num int) (int, error) {
    if num < 0 {
        return -1, fmt.Errorf("wow~")
    }
    return num, nil
}

func panicDie() {
    // 没办法的时候才这么玩
    panic("Game Over!")
}

func main() {
    useErr()
    newError()
    fmtError()
    panicDie() // program finished
    fmt.Println("You can not see this.")
}