package main
import "fmt"
//import "unsafe"


// https://www.runoob.com/go/go-structures.html

func test(var1 string) int {
    var balance [10] float32
    var balance2 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    var balance3 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    
    var salary float32 = balance[9]

    fmt.Println("begin test------------")     
    fmt.Println(salary) // 0
    fmt.Println(balance[1]) // 0
    fmt.Println(balance2[2]) // 3.4
    fmt.Println(balance3[3]) // 7
    fmt.Println("end test------------")
    return len(var1)
}

func pointer() {
    var ptr *int
    ptr = nil
    fmt.Println("pointer", ptr)
}

func swap(x, y string) (string, string) {
   return y, x
}


func swap2(x *int, y *int) {
   var temp int
   temp = *x
   *x = *y
   *y = temp
}


func main() {
    // var bb bool = true
    //var age int
    var a, b int = 1, 2
    var s string = "str"
    var ar []int
    var e error
    var f func(string) int
    var ch chan int
    const c string = "Im const var"
    const defc = "abc"
    const (
        unknown = 0
        female = 1
        male = 2
    )
    const (
        a1 = iota
        b1 = iota
        c1 = iota
    )
    const (
        a2 = iota
        b2
        c2
    )
    
    newvar := 11 // 等效于 var newvar = 11    
    a, b, s = 3, 4, "hia hia hia"
    a = len(s)
    //b = unsafe.Sizeof(male)
    
    fmt.Println("Hello" + "World")
    fmt.Println(a, b)
    fmt.Printf("%v %v %v %q\n", s, f, e, ch)
    fmt.Printf("%v\n", ar)
    fmt.Println(&newvar)
    fmt.Println(b1, b2)
    
    i := 0
//     for i := 0; i < len(a); i++ {
    for true {
        if i > 5 {
            break
        }
        fmt.Println(i)
        i = i + 1
    }
    
    fmt.Println(test("test"))
    
    s1, s2 := swap("Google", "baidu")
    fmt.Println(s1, s2)
    
    swap2(&a, &b)
    fmt.Println(a, b)

    pointer()
}



