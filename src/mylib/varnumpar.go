package mylib

import (
	"fmt"
)

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	// _丢弃掉索引
	// 第一个是索引，第二个是具体的值，类似于C#的foreach
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func Typecheck(values ...interface{}) {
	for _, value := range values {
		switch v := value.(type) {
		case int:
			fmt.Printf("这个是一个int的参数%v", v)
		default:
			fmt.Print("这个不是一个int的参数")
		}
	}
}
