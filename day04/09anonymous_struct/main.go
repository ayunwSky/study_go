package main

import "fmt"

// 结构体的匿名字段

// 相同类型只能写一个，即 结构体中字段是唯一，只能写一个string，一个 int。
// 结构体的匿名字段类型不常用，主要用于字段少且字段比较简单的场景
type person struct {
	string
	int
}

func main() {
	p1 := person{
		"艾伦",
		18,
	}

	fmt.Println(p1)
	// 通过类型当做名称来进行获取里面的value
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
