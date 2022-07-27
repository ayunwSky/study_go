package main

import "fmt"

/*
	Go 中的return不是原子操作，底层是分为两个步骤执行：
		第一步：返回值赋值
		第二步：真正的 RET 返回

	函数中如果存在defer，则defer执行的时间是在第一步和第二步之间，即：
		第一步：返回值赋值
		执行 defer 语句
		第二步：真正的 RET 返回
*/

func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是 x 不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值=x
}

func f3() (y int) {
	x := 5
	defer func(x int) {
		x++	// 修改了 x
	}(x)
	return x	// 返回值 y=x=5,
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5 // 返回值 = x = 5
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
