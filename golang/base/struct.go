package main
import "fmt"
import "unsafe"

// https://www.runoob.com/go/go-structures.html

type Book struct {
    title string
    id int
    isbn int
}

func call(pbook *Book) {
    fmt.Printf("func_call:%d\n", pbook.id)
    pbook.id = -1
}

func main() {
    var book Book
    var pbook *Book = nil

    book.title = "Go"
    book.id = 1

    pbook = &book

    fmt.Printf("size=%d\n", unsafe.Sizeof(book)) // 32
    fmt.Printf("size=%d\n", unsafe.Sizeof(book.id)) //8
    fmt.Println(pbook.id)
    fmt.Println(Book{title:"Go", id:2})

    call(pbook) // 1
    call(&book) // -1
}
