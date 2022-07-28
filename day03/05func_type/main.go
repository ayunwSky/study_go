package main

import "fmt"

// 函数类型作为参数和返回值

func f1() {
	fmt.Println("I am f1 func")
}

func f2() int {
	return 10
}

// 函数作为参数，且有一个 int 类型的返回值
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func f4(x, y int) int {
	return x + y
}

// 函数作为返回值
func f5(x func() int) func(int, int) int {
	ret := func(a, b int) int {
		return a + b
	}
	return ret
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	fmt.Println(a)

	b := f2
	fmt.Printf("%T\n", b)

	f3(f2)

	fmt.Printf("%T\n", f4)

	f5Res := f5(f2)
	fmt.Println(f5Res)
	fmt.Printf("%T\n", f5Res)
}
