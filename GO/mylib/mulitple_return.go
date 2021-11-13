package mylib

func getX2AndX3_2(input int) (x1 int, x2 int) {
	x1 = 0
	x2 = 1
	return
}

// c表示的是一个指针，因为是引用传递，它将会改变c传入进来的那个参数值，使用这个方法，建议说明清楚
func Multiply(a, b int, c *int) {
	*c = a * b
}
