package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readAll(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	fmt.Println(string(content))
}

func readByte(filename string) {
	fi, err := os.Open(filename)
	if err != nil {
	   panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)
 
	chunks := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
	   n, err := r.Read(buf)
	   if err != nil && err != io.EOF {
		  panic(err)
	   }
	   if 0 == n {
		  break
	   }
	   fmt.Println(string(buf[:n]))
	   chunks = append(chunks, buf[:n]...)
	   break
	}
	fmt.Println(string(chunks))
}

func main() {
	readAll("test.log")
	readByte("test.log")
}
