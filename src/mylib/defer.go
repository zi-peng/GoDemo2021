package mylib

import "fmt"

func Function1() {
	fmt.Printf("你好啊!\n")
	//效果类似于C#中的finally,在函数执行完之后进这里。
	defer function2()
	fmt.Printf("好久不见!\n")
}
func function2() {
	fmt.Printf("小明\n")
}

//defer的位置很重要，因为它的值一旦确定不会发生改变,
//defer 必须放置return前使用，后面使用不会被执行
func Function3() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

//defer 遵循后进先出的原则，跟栈一样，当一个方法中出现多个defer的时候需要注意
func Function4() {
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d", i)
	}
}
