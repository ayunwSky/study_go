package main

import "fmt"

/*
	切片是一个引用类型，包含地址、长度和容量
*/

func main() {
	// 切片的定义
	var s1 []int    // 定义一个存放 int 类元素的切片
	var s2 []string // 定义一个存放 string 类元素的切片
	fmt.Println(s1, s2)

	// nil 是空的意思
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"上海", "杭州", "北京"}
	fmt.Println(s1, s2)

	// 上面已经初始化过 s1 和 s2，所以这里不是空，则结果为 false
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	// 长度和容量
	fmt.Printf("len(s1): %d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2): %d cap(s2):%d\n", len(s2), cap(s2))

	// 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	// 基于a1数组切片，从第 0 个索引开始取到第 3 个索引，左闭右开，右边的 4 不包含，只能取到 3
	s3 := a1[0:4]
	fmt.Println(s3) //[1 3 5 7]

	s4 := a1[1:6]
	fmt.Println(s4) // [3 5 7 9 11]

	// 省略开头相当于从索引 0 开始切到索引 4
	s5 := a1[:4] // a1[0:4] ---> [1 3 5 7]
	// 省略几位相当于从第三个索引开始切到最后一个，包含最后一个
	s6 := a1[3:] // a1[3:len(a1)]  ---> [7 9 11 13]
	// 相当于从索引 0 开始切到最后一个，包含左后一个
	s7 := a1[:] // a1[0:len(a1)] ---> [1 3 5 7 9 11 13]

	fmt.Println(s5)
	fmt.Println(s6)
	fmt.Println(s7)

	/*
		切片指向了一个底层的数组
		切片的长度就是它的元素个数
		切片的容量指的是底层数组的容量，而底层数组是从切片的第一个元素开始算，一直到最后一个元素

		参考文中切片的本质的图示: https://www.liwenzhou.com/posts/Go/06_slice/
	*/

	// 底层数组 a1 的长度是7，切片 s5 是从 a1 的第 0 个索引开始算，所以这里 cap 就是从第 0 个元素一直到最后一个元素也就是7
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5)) // len(s5):4 cap(s5):7
	// 底层数组 a1 的长度是7，切片 s6 是从 a1 的第 3 个索引开始算，一直到最后一个索引，所以这里的 cap 是 4
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6)) // len(s6):4 cap(s6):4

	// 切片再切片
	s8 := s6[3:]
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8)) // len(s8):1 cap(s8):1

	// 切片是引用类型，都指向了底层的数组
	// 更改了底层数组后，切片的值就变成了元素被更改后的值
	fmt.Println("s6:", s6) // s6: [7 9 11 13]

	a1[6] = 1300           // 修改了底层数组第6个元素的值
	fmt.Println("s6:", s6) // s6: [7 9 11 1300]
	fmt.Println("s8:", s8) // s8: [1300]

}
