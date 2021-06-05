package main
import "fmt"
import "unsafe"

func useSlice() {

    var items = [...] int{1, 2, 3, 4, 5}
    // make = create a slice
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

    fmt.Printf("useSlice:%d\n", slice4[0])
    fmt.Printf("useSlice:%d\n", unsafe.Sizeof(slice1))
    fmt.Printf("useSlice:%d\n", unsafe.Sizeof(slice5))
    for i := 0;i<len(slice5);i++ {
        fmt.Printf("\t%d ", slice5[i])
    }
    items[3] = -items[3]
    fmt.Printf("\nuseSlice:5=%d %d\n", len(slice5), cap(slice5)) 
    for i := 0;i<len(slice6);i++ {
        fmt.Printf("useSlice:%d ", slice6[i])
    }
    fmt.Printf("\nuseSlice:%d\n", len(slice6)) // 2
    fmt.Printf("useSlice:%d %d\n", len(slice3), cap(slice3)) // 2 3
}
func useAdd() ([]string) {
    var s1 []string = make([]string, 2)
    s1[0] = "A"
    s1[1] = "B"
    s2 := append(s1, "C", "D") // [A,B,C,D], len==4
    fmt.Println("useAdd:", s2, len(s2))
    s3 := append(s2[:2], s2[2+1:]...) // remove C
    fmt.Println("useAdd:", s3, len(s3)) // [A B D] len==3

    return s3
}

func useCopy() {
    var dest = make([]string, 2)
    var src = useAdd()
    copy(dest, src) // [A B D] --> [A B]
    fmt.Println("copy: ", src, dest)
    src[0] = "X"
    fmt.Println("copy: ", src, dest) // [X B D] [A B]
}

func useMap()  {
    var p1 = make(map[string] int) // string key, int value
    p1["tom"] = 27
    p1["lake"] = 61
    fmt.Println("useMap:", p1["tom"])
    p1["some"] = 13
    delete(p1, "tom") // only use for map(not use for slice)
    fmt.Println("useMap:", p1)
}

func main() {
    useSlice()
    useAdd()
    useCopy()
    useMap()
}