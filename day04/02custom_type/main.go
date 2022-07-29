package main

import "fmt"

// 自定义类型和类型别名

// type 后面跟的是类型

// 自定义类型
// 以下的意思是 基于一个内置的 int 类型，造了一个 myInt 的类型
type myInt int

// 类型别名
type yourInt = int

func main() {
	var n myInt
	n = 100
	fmt.Printf("%T\n",n)	// main.myInt
	fmt.Println(n)

	fmt.Println()

	var m yourInt
	m = 20
	fmt.Printf("%T\n",m)	// int
	fmt.Println(m)

	/*
		n的类型是：main.MyInt，表示main包下定义的 MyInt 类型
		m的类型是：int
		
		MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型
	*/
}
