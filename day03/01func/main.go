package main

import "fmt"

func f1() {
	fmt.Println("hello 浙江")
}

func f2(name string) {
	fmt.Println("你好", name)
}

func f3(x int, y int) int {
	sum := x + y
	return sum
}

func f4(x, y int) int {
	return x + y
}

func f5(title string, y ...int) int {
	fmt.Println(title)
	fmt.Println(y) // y 是 int 类型的切片
	return 1
}

func f6(x, y int) (sum int) {
	sum = x + y
	// fmt.Println("sum:", sum)
	return
}

func f7(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

func main() {
	f1()
	f2("小姜")

	f3Res := f3(10, 20)
	fmt.Println(f3Res)

	f5("老王", 1, 2, 3, 4, 5)

	f6Res := f6(100, 200)
	fmt.Println(f6Res)

	f7ResSum, f7ResSub := f7(100, 20)
	fmt.Println(f7ResSum, f7ResSub)
}
