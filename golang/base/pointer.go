package main
import "fmt"

// https://www.runoob.com/go/go-pointers.html
func main() {
    var a int = 10
    var b *int
    var ptr *int

    b = &a

    fmt.Printf("a=%d\n", a)
    fmt.Printf("&a=%x\n", &a)
    fmt.Printf("b=%x\n", b)
    fmt.Printf("*b=%d\n", *b)
    fmt.Printf("null ptr=%x\n", ptr)

    if (nil == ptr) {
        fmt.Printf("ptr is null\n")
    }
}