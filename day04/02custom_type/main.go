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
	fmt.Printf("%T\n",n)
	fmt.Println(n)

	fmt.Println()

	var m yourInt
	m = 20
	fmt.Printf("%T\n",m)
	fmt.Println(m)
}
