package main
import "fmt"
import "unsafe"

// https://www.runoob.com/go/go-slice.html

func main() {
    var items = [...] int{1, 2, 3, 4, 5}
    // create
    var slice1 []int = make([]int, 3)
    var slice2 []int

    slice3 := make([]int, 2, 3)

    // initialize
    slice4 := [] int{1,2}
    slice5 := items[:]

    slice6 := items[2:4] // val=3, 4

    if (nil == slice2) {
        fmt.Printf("slice2 is nil\n")
    }

    fmt.Printf("%d\n", slice4[0])
    fmt.Printf("%d\n", unsafe.Sizeof(slice1))
    fmt.Printf("%d\n", unsafe.Sizeof(slice5))
    for i := 0;i<len(slice5);i++ {
        fmt.Printf("%d ", slice5[i])
    }
    items[3] = -items[3]
    fmt.Printf("\n5=%d %d\n", len(slice5), cap(slice5)) 
    for i := 0;i<len(slice6);i++ {
        fmt.Printf("%d ", slice6[i])
    }
    fmt.Printf("\n%d\n", len(slice6)) // 2
    fmt.Printf("%d %d\n", len(slice3), cap(slice3)) // 2 3
}