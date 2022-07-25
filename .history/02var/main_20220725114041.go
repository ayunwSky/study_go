package main

import "fmt"

批量声明
var (
	name string = "allen"
	age  int    = 18
	isOk bool   = true
)

func foo() (int, string) {
	return 10, "steve"
}

func main() {
	// 简短变量声明。只能在函数中声明
	m := 20
	n := 10

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isOk)

	fmt.Println()

	fmt.Println("m=",m)
	fmt.Println("n=",n)
	
	fmt.Println()

	// 匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明
	x, _ := foo()
	_, y := foo()

	fmt.Println("x=",x)
	fmt.Println("y=",y)
}
