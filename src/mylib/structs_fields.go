package mylib

import "fmt"

type A struct {
	Num int
	B   float32
	C   string
}

func StructDemo(a A) A {
	// fmt.Sprintf("a的值是%d", a.Num)
	a.Num = 200
	fmt.Printf("a的值是%d", a.Num)
	return a
}
