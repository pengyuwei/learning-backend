package main

import (
	"errors"
	"fmt"
)

type Robot interface {
	PowerOn() error
}

// 实现接口的方法集，并保证函数签名正确即可
type HAL9000 struct {
	Name string
}

func (a *HAL9000) PowerOn() error {
	return nil
}

type R2D2 struct {
	Broken bool
}

func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken") //
	} else {
		return nil
	}
}

// 接口也是类型，可以作为参数传递
func Boot(r Robot) error {
	return r.PowerOn()
}

func main() {
	h := HAL9000{
		Name: "HAL9000",
	}

	r := R2D2{
		Broken: true,
	}

	err := Boot(&r)
	if err == nil {
		fmt.Println("OK!")
	} else {
		fmt.Println(err)
	}

	err = Boot(&h)
	if err == nil {
		fmt.Println("HAL is powered on!")
	} else {
		fmt.Println(err)
	}
}
