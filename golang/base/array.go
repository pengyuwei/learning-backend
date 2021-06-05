package main
import "fmt"

// https://www.runoob.com/go/go-arrays.html

func array()  {
    var games [3]string
    games[0] = "pal"
    games[1] = "tank"
    games[2] = "quake"
    fmt.Println(games) // [pal tank quake]
}

func main() {
    var items[10] float32
    var items2 = [5] float32{100.0, 2.0, 3.14, 7.0, 20}
    var items3 = [...] float32{100.0, 1.0}
    items3[1] = 2.0

    var a float32 = items2[2] // 3.14

    fmt.Printf("a=%f, item3=%f, items1[0]=%f\n", a, items3[1], items[0])

    array()
}