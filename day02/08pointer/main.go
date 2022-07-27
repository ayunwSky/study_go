package main

import "fmt"

// vscode 国内不支持 go module ，开启后会特别卡
/*
	指针 pointer
		& 取内存地址
		* 根据内存地址取值
*/

func main() {
	// // 1. & 取内存地址
	// n := 18
	// fmt.Println(&n)

	// // 2.* 根据内存地址取值
	// ptr := &n
	// fmt.Println(*ptr)
	// fmt.Printf("%T\n",ptr)	// *int 表示 int类型的指针，就是一个内存地址

	// m := *ptr
	// fmt.Println(m)

	fmt.Println()

	// var a *int	// nil,a是一个空的指针，没有初始化，没有内存地址，所以会报错
	// // *a = 100
	// fmt.Println(*a)

	// new 函数申请（开辟）一个内存地址
	var a *int
	fmt.Println(a) // 没有赋予初始值，因此是一个 nil

	var b = new(int) // new 开辟了新的内存地址
	fmt.Println(*b)  // 内存地址的 0 值，
	*b = 100         // 赋予了初始值 100
	fmt.Println(*b)  // 100

	// make 用于分配内存，和new的区别是：make只用于 slice、map以及channel

}
