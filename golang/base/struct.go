package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Book struct {
	title string
	id    int
	isbn  int
}

func call(pbook *Book) {
	fmt.Printf("func_call:%d\n", pbook.id)
	pbook.id = -1
}

func useStruct() {
	type Movie struct {
		Name   string
		Rating float32
	}
	var m1 Movie
	m1.Name = "3 idiot"
	m1.Rating = 9.9
	fmt.Println("useStruct:", m1)

	m2 := Movie{
		Name:   "Tron",
		Rating: 9.8, // 分行写的时候，最后一行必须有逗号
	}
	fmt.Println("useStruct:", m2.Name, m2.Rating)

	m3 := new(Movie)
	m3.Name = "Contact"
	m3.Rating = 10
	fmt.Printf("useStruct: %+v\n", m3)

	m4 := Movie{Name: "A man from earth", Rating: 9.7}
	m5 := Movie{"A man from earth", 9.7}
	fmt.Printf("useStruct: %+v\nuseStruct: %v\n", m4, m5)

	if m4 == m5 {
		fmt.Printf("useStruct: same\n")
	}

	fmt.Println("useStruct: type=", reflect.TypeOf(m1)) // main.Movie

	clone := Movie{}
	clone = m5 // clone, not refrence
	clone.Name = "CloneMan"
	fmt.Println("useStruct: ", m5, clone)

	c2 := &m5 // point
	c2.Rating = 6
	fmt.Println("useStruct: ", m5, *c2) // same

}

func useAdvanceStruct() {
	type Address struct {
		City   string
		Street string
	}
	type Superhero struct {
		Name    string
		Address Address
	}
	e := Superhero{
		Name: "Yang",
		Address: Address{
			City:   "Beijing",
			Street: "Zhongguancun",
		},
	}
	fmt.Printf("%+v\n", e)
}

func main() {
	var book Book
	var pbook *Book = nil

	book.title = "Go"
	book.id = 1

	pbook = &book

	fmt.Printf("size=%d\n", unsafe.Sizeof(book))    // 32
	fmt.Printf("size=%d\n", unsafe.Sizeof(book.id)) //8
	fmt.Println(pbook.id)
	fmt.Println(Book{title: "Go", id: 2})

	call(pbook) // 1
	call(&book) // -1

	useStruct()
	useAdvanceStruct()
}
