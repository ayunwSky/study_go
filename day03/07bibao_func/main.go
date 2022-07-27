package main

import "fmt"

// 闭包函数

func f1(f func()) {
	fmt.Println("f1 func...")
	f()
}

func f2(x, y int) {
	fmt.Println("f2 func...")
	fmt.Println(x + y)
}

// 定义一个函数对f2进行包装
func f3() {

}

func main() {

}
