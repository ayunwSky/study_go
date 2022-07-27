package main

import (
	"fmt"
)

// 函数：函数在Go语言中属于一等功明
/*
	Go 语言中支持 函数、匿名函数和闭包

	函数定义：
	func 函数名(参数)(返回值){
    	函数体
	}
*/

// 没有参数，没有返回值的函数
func f1() {
	fmt.Println("f1():没参数，没返回值的函数")
}

// 有参数，没有返回值
func f2(a int, b int) {
	fmt.Println("f2()没返回值:", a+b)
}

// 没有参数但是有返回值
func f3() int {
	return 3
}

// 返回值可以命名也可以不命名
// 命名返回值就相当于在函数中声明了这个变量
// 返回值命名后，在return返回的时候，可以写成 return ，也可以写成 return ret
func f4(x, y int) (ret int) {
	ret = x + y
	fmt.Println("f4() ret:", ret)
	return ret // 使用命名返回值，return后的返回值可以省略，也可以写上
}

// 多个返回值
func f5() (int, string) {
	return 1, "杭州"
}

// 参数的类型简写
// 当参数中连续两个参数类型一致时，可以将非最后一个参数类型省略
func f6(x, y int, m, n string, i, j bool) int {
	return x + y
}

// 可变长参数
// 可变长参数必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)	// 调用函数 f7 的时候，y的值可以不传，y的类型是切片 []int
}

func sayHello() {
	fmt.Println("sayHello(): 小姜呀!")
}

// Go语言中 函数没有默认参数的概念

// 有参数，有返回值
// 同类型的x,y，则可以直接在y后面指定一次参数类型即可
func intSum(x, y int) int {
	return x + y
}

func main() {
	f1()
	f2(1, 2)
	f3()
	f4(50, 100)

	_, n := f5()
	fmt.Println("f5() n:", n)

	f7("下雨了")
	f7("下雨了", 0, 1, 2, 3, 4, 5)

	sayHello()
	res := intSum(10, 20)
	fmt.Println(res)
}
