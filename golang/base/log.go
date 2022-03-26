package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    log.Fatal(err) // finish program. output format: 2022/03/19 22:45:09 <nil>
    fmt.Println(input, "End.") // not display
}