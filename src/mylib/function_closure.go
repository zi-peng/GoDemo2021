package mylib

func Adder() func(int) int {
	var x int //这个值在被调用时会被存储，因为他是先声明的
	return func(i int) int {
		x += i
		return x
	}
}
