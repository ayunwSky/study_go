package main

import "fmt"

// 匿名函数
/*
	函数内部无法定义带名字的函数，因此在函数内部就只能定义匿名函数
	如果函数只是被调用一次，则可以简写成立即执行函数,就是写完函数后加括号立即执行
*/

// var f1 = func(x, y int) {
// 	fmt.Println(x + y)
// }

func main() {
	// f1(10,20)
	// 匿名函数
	// 函数内部无法声明带名字的函数，所以匿名函数一般用在函数内部
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)

	// 函数只是被调用一次，则可以简写成立即执行函数,就是写完函数后加括号立即执行
	// func(){
	// 	fmt.Println("hello world")
	// }()

	func(x, y int) {
		fmt.Println(x + y)
		fmt.Println("hello world")
	}(100, 200)

}
