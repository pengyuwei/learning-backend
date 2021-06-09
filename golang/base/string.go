package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func useStr() {
    s := `AI
    Artificial intelligencr`
    fmt.Println(s)
    /*
       %q: 'A'
       %c: A
       %v: 65
       %b: 1000001
       %d: 65
    */
    fmt.Printf("useStr:%d\n", s[0]) // 65=A

    s = "Hello" + " world "
    s = strings.TrimSpace(s)
    s += "!\n"
    var str string = s + "... "
    str = strings.ToLower(str)
    fmt.Printf("useStr:%d,%v\n", len(str), str)

    var ret int = strings.Index(str, "world")
    if ret == -1 {
        fmt.Println("useStr: Not found")
    } else {
        fmt.Println("useStr:", ret)
    }

}

func conv() {
    var i int = 1
    strI := strconv.Itoa(i)
    fmt.Println("conv:", strI)
}

func buff() {
    // 性能好于+=（30倍）
    var buffer bytes.Buffer
    for i := 0; i < 50; i++ {
        buffer.WriteString(".")
    }
    fmt.Println("buff:", buffer.String())
}

func main() {
    useStr()
    conv()
    buff()
}
