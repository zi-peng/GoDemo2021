package mylib

import "fmt"

func Literal() {
	defer func() {
		fmt.Print("defer，函数调用完成之前调用这个")
	}()
}
