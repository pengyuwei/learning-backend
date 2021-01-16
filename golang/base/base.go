package main
import "fmt"
import "reflect"
import "strconv"
//import "unsafe"


var global int = 1

func scope() {
    fmt.Println("-----------scope begin----------------------")
    fmt.Println(global) // 1
    {
         global := 2
        fmt.Println(global) // 2
    }
    fmt.Println(global) // 1

    var global int = 3
    fmt.Println(global) // 3
}

func test(var1 string) int {
    var balance [10] float32
    var balance2 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    var balance3 = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    
    var salary float32 = balance[9]

    fmt.Println("-----------test begin------------")
    fmt.Println(salary) // 0
    fmt.Println(balance[1]) // 0
    fmt.Println(balance2[2]) // 3.4
    fmt.Println(balance3[3]) // 7
    return len(var1)
}

func pointer() {
    fmt.Println("-----------pointer begin------------")
    var ptr *int
    ptr = nil
    fmt.Println("pointer", ptr)

    sp := "Hello world"
    fmt.Println("&sp=", &sp)
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

func tryvar()  {
    fmt.Println("-----------tryvar begin------------")
    var a, b int = 1, 2
    var s string = "str"
    var z string
    if z == "" {
        fmt.Println("z is zero") // TODO
    }
    a, b, s = 3, 4, "hia hia hia"
    a = len(s)
    s += "haha"
    fmt.Println(a, b)
    fmt.Printf("%v\n", s)

    var f float32 = 3.14159265358979
    fmt.Println(f)

    var bb bool = true
    fmt.Println(bb)

    var ar []int
    fmt.Printf("%v\n", ar)

    var e error
    fmt.Printf("%v\n", e)

    const c string = "Im const var"
    newvar := 11 // 等效于 var newvar = 11    
    fmt.Println(&newvar)
    
    //b = unsafe.Sizeof(male)
    fmt.Println("Hello" + "World")

    s1, s2 := swap("Google", "baidu")
    fmt.Println(s1, s2)
    
    swap2(&a, &b)
    fmt.Println(a, b)

}

func typeof()  {
    var s string = "dim"
    var i int = 11
    var f float32 = 6.18

    fmt.Println(reflect.TypeOf(s))
    fmt.Println(reflect.TypeOf(i))
    fmt.Println(reflect.TypeOf(f))
}

func convert()  {
    var s string = strconv.FormatBool(true)
    fmt.Println("-----------convert begin-----------")
    fmt.Println(reflect.TypeOf(s), s)
}

func constvar() {
    fmt.Println("-----------constvar begin-----------")
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

    fmt.Println(b1, b2)
}

func main() {
    var f func(string) int
    var ch chan int

    fmt.Printf("%v %q\n", f, ch)
    
    i := 0
    for true {
        if i > 5 {
            break
        }
        fmt.Println(i)
        i = i + 1
    }
    
    fmt.Println(test("test"))

    constvar()
    pointer()
    tryvar()
    typeof()
    convert()
    scope()
}



