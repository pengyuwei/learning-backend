/*
Package shows how to use godoc.
And this is second row.
git clone https://github.com/shapeshed/golang-book-examples.git
cd golang-book-examples/hour14/
go doc example03.go
*/
package docexample

import "fmt"

// Demo 显示使用方法.
func demo() {
    fmt.Println("以要解释名称的大写字母开头，以点结束")
    fmt.Println("go doc main.go")
}

/* 
Main shows how to use godoc.
And this is second row.
*/
func main() {
    demo()
}
