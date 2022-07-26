package main

import (
	"fmt"
)

/*
 	在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化
	数组定义: var 数组变量名 [元素数量]T，如：var a [3]int
	数组的长度是数组类型的一部分，即类型包含了长度，var a1 [3]bool 和 var a2 [4]bool
	是不同的类型
*/

func main() {
	var a1 [3]bool // [false false false]
	var a2 [4]bool // [false false false false]
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	// 数组初始化
	// 数组不初始化：默认元素都是零值(布尔值:false，整型和浮点型都是0,string: "",)
	fmt.Println(a1, a2)

	fmt.Println()

	// 初始化方法1
	a1 = [3]bool{true, false, true}
	fmt.Println(a1)

	// 初始化方法2
	// a10 := [10]int{0,1,2,3,4,5,6,7,8,9}

	fmt.Println()

	// 根据初始值自动推断数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%d\n", a10)

	fmt.Println()

	// 初始化方法3
	// a3 :=[5]int{1,2}	// 默认除了1和2，后面都是0值。[1 2 0 0 0]

	// 根据索引初始化
	a3 := [5]int{0: 1, 4: 3}
	fmt.Println(a3)

	// 数组遍历
	cities := [...]string{"北京", "上海", "广州", "深圳", "杭州"}
	// 方法1:根据索引遍历
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	fmt.Println()

	// 方法2:for range
	for k, v := range cities {
		fmt.Println(k, v)
	}

	fmt.Println()

	// 多维数组,默认数组中只能放同一种类型
	// 创建一个二维数组：[[1 2] [3 4] [5 6]]
	var a11 [3][2]int
	// 初始化
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	fmt.Println()

	// 多维数组遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	fmt.Println()
	// 数组是值类型
	/*
		相当于 ctrl+c,ctrl+v拷贝了一份一模一样的数据
		更改原来的数据，拷贝过去的数据不会受到影响。

		可以理解为将目录A下的一份world文档拷贝到目录B下
		更改目录A下的world，不会影响目录B下的world

		如果是引用类型，则相当于linux下的软链接，
		则更改原来的数据，拷贝过去的数据也会受到影响
	*/
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Printf("b1=%d\n", b1)
	fmt.Printf("b2=%d\n", b2)

	fmt.Println()
	// 练习题
	/*
		1、求 数组 [1,3,5,7,8] 所有元素的和
		2、找出数组中和为指定值的两个元素的下标，比如从数组[1,3,5,7,8]中找出 和为8的两个元素的下标
			分别为(0,3) 和 (1, 2)
	*/

	c1 := [5]int{1, 3, 5, 7, 8}

	// 题目1
	sum := 0
	for _, v := range c1 {
		sum += v
	}
	fmt.Println(sum)

	fmt.Println()
	// 题目2
	// 定义两个for循环，外层从第1个元素开始遍历，内层循环从外层循环的后面一个开始循环
	for i := 0; i < len(c1); i++ {
		for j := i + 1; j < len(c1); j++ {
			if c1[i]+c1[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}

}
