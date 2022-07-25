package main

import "fmt"

// 1、变量
// var name string = "allenjol"
// var age int = 18

// 批量声明变量
// 在全局声明的变量，如果没有使用，编译可以通过
// 在函数中声明的变量，如果没使用，编译会报错
var (
	name string = "allen"
	age  int    = 18
	isOk bool   = true
)

func foo() (int, string) {
	return 10, "steve"
}

// 2、常量
// const pi = 3.1415926
// const e = 2.7182

// 批量声明常量
// const (
// 	pi = 3.1415926
// 	e  = 2.7182
// )

// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
const (
	n1 = 100
	n2
	n3
)

func main() {
	// 简短变量声明。只能在函数中声明
	m := 20
	n := 10

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isOk)

	fmt.Println()

	fmt.Println("m=", m)
	fmt.Println("n=", n)

	fmt.Println()

	// 匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明
	x, _ := foo()
	_, y := foo()

	fmt.Println("x=", x)
	fmt.Println("y=", y)
}
