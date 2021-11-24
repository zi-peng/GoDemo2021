package mylib

import "fmt"

func Mapdemo() {
	//声明一个map
	//如果想给这个map赋值，得先初始化
	m1 := make(map[string]int, 10)
	m1["理想"] = 100
	//判断map中是否有个某个值
	v, ok := m1["理想1"]
	if ok {
		fmt.Printf("%v", v)
	} else {
		fmt.Printf("你查询的值不存在!")
	}
	fmt.Printf("m1的值是%v", m1["理想"])
	//遍历Key
	for k, v := range m1 {
		fmt.Printf("key的值是%v，v的值是%v", k, v)
	}
	//删除指定key的map
	delete(m1, "理想")
}
