package main

import (
	"fmt"
	"sort"
)

// copy函数复制切片

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1 // 赋值
	var a3 = make([]int, 3, 3)
	copy(a3, a1) // 使用 copy() 函数将切片 a1 中的元素 复制到切片 a3 中

	fmt.Println(a1, a2, a3)

	a1[0] = 100
	fmt.Println(a1, a2, a3)

	// 从切片中删除元素
	a5 := []int{10, 20, 30, 40, 50, 60}
	// 删除索引为 2 的元素,也就是 30
	// Go 中没有直接删除切片中元素的方法，可以通过 append() 函数 来实现 删除的功能
	a5 = append(a5[:2], a5[3:]...)
	fmt.Println(a5) // [10 20 40 50 60]

	fmt.Println()


	// 练习题,面试题
	var m = make([]int, 5, 10)	// 声明一个被 make 函数 初始化过的切片 m
	fmt.Println(m)	// [0 0 0 0 0]

	for i := 0; i < 10; i++ {
		m = append(m, i)
	}
	fmt.Println(m)	// [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	fmt.Println(cap(m))

	fmt.Println()

	// 对切片进行排序
	var k = [...]int{3,7,9,2,0,8}
	fmt.Println("排序前的切片k:",k)
	sort.Ints(k[:])
	fmt.Println("排序后的切片k:",k)	// [0 2 3 7 8 9]

	/*
		切片总结：
			要从切片a中删除索引为index的元素，操作方法是 a = append(a[:index], a[index+1:]...)
	*/
}
