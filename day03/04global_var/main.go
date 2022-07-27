package main

import (
	"fmt"
)

// 变量作用域

var x = 100

func f1() {
	// name := "姜"
	// 函数中查找变量的顺序
	/*
		1. 先在函数内部查找
		2. 如果函数内部找不到就去函数外部查找，比如外层函数内部
		3. 如果外层函数没有就去全局找，全局都没有就报错了，比如提示：undefined
	*/
	fmt.Println(x)
}
func main() {
	f1()
	// fmt.Println(name)	// f1 函数内部的 name 变量只能在 f1 函数内使用

	if i := 10; i < 18 {
		fmt.Println("去上学")
	}
	// fmt.Println(i)	// 只能在 if 语句块中执行
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
	// fmt.Println(j)	// 只能在 for 语句块中执行
}
