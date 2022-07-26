package main

import "fmt"

// name 函数创建切片

func main() {
	// 声明一个int类型的切片，长度为 5，容量不写默认和长度一致
	s1 := make([]int, 5)                                              // s1 := make([]int, 5, 10) 表示长度 5，容量 10
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1)) // s1=[0 0 0 0 0] len(s1)=5 cap(s1)=5

	// 声明一个int类型的切片，长度为 0，容量 10
	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2)) // s2=[] len(s2)=0 cap(s2)=10

	/*
		切片不能直接比较，不能用 == 操作符来判断两个切片是否含有全部相等元素，切片唯一合法的
		比较操作是和 nil 进行比较。
		一个 nil 值的切片并没有底层数组，一个 nil 值的切片的长度和容量都是 0
		但是我们不能说一个长度和容量都是 0 的切片一定是 nil

		所以要判断一个切片是否为空，需要用 len(s) == 0 来判断，而不能用 s == nil 来判断

		比如下面例子
	*/
	fmt.Println()

	var s10 []int // len(s10)=0,cap(s10)=0,s10 = nil
	fmt.Printf("s10=%v len(s10)=%d cap(s10)=%d\n", s10, len(s10), cap(s10))

	s11 := []int{} // len(s11)=0,cap(s11)=0,s11 ！= nil
	fmt.Printf("s11=%v len(s11)=%d cap(s11)=%d\n", s11, len(s11), cap(s11))

	s12 := make([]int, 0)
	fmt.Printf("s12=%v len(s12)=%d cap(s12)=%d\n", s12, len(s12), cap(s12))

	fmt.Println()

	// 切片的赋值拷贝
	s15 := []int{1, 3, 5}
	s16 := s15 // s15 和 s16 都指向了同一个底层数组
	fmt.Println(s15, s16)

	s15[0] = 100
	fmt.Println(s15, s16)

	fmt.Println()

	// 切片的遍历
	s17 := []int{1, 3, 5}
	// 1、索引遍历
	for i := 0; i < len(s17); i++ {
		fmt.Println(s17[i])
	}

	// 2、for range 遍历
	for k, v := range s17 {
		fmt.Println(k, v)
	}
}
