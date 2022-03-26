package main

import "fmt"

// 一般情况，第二个返回参数用于返回错误信息
// 返回多个值，需要标明返回值类型
func RetMultiVals() (int, string)  {
    return 1, "string"
}

func callPublicFunc() {
    // 对于不打算使用的变量，可以使用下划线(空白标识符)代替，就不会引发编译错误了.
    a, _ := RetMultiVals()
    fmt.Println("callPublicFunc", a)
}

// 无限参数的函数
func sum(args...int) int {
    // args is slice
    total := 0
    for _, itor := range args {
        total += itor
    }
    return total
}

// 指定返回值变量名的函数，会降低代码可读性
func fnWithName() (x,y string) {
    x = "RetVal.1"
    y = "RetVal.2"
    return // naked return
}

// 参数为函数
func fnWithFunc(f func() string) string {
    return f()
}

func fnWithDefer() {
    defer fmt.Println("defer fn3") // output when function finished
    defer fmt.Println("defer fn2") // output when function finished
    fmt.Println("fnWithDefer1") // first output
}

func main() {
    var ret int

    ret = sum(1,2,3,4,5)
    fmt.Println("sum=", ret)

    fmt.Println(fnWithName())

    fn := func ()  {
        fmt.Println("Functional Programming")
    }
    fn()
    fn2 := fnWithName
    fmt.Println(fn2())

    fn3 := func() string {
        return "fnWithFunc"
    }
    fmt.Println(fnWithFunc(fn3))

    fnWithDefer()

    ret1, ret2 := RetMultiVals()
    fmt.Println(ret1, ret2)

    callPublicFunc()
}