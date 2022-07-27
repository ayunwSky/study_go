package main

import "fmt"

/*
	append() 函数 为切片追加元素
	copy() 函数 复制切片
*/

func main() {
	s1 := []string{"北京", "上海", "杭州"}
	fmt.Printf("切片追加元素前: s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	// s1[3] = "广州"	// 错误的写法，会导致编译错误，索引越界（index out of range）

	// 调用 append() 函数追加元素，必须用原来的切片变量接收返回值
	// append() 函数追加元素，原来的底层数组放不下的时候,Go 底层就会把底层数组换一个新的大的底层数组
	// 必须用变量接收 append 的返回值
	s1 = append(s1, "广州")
	fmt.Printf("切片第一次追加元素后: s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	s1 = append(s1, "深圳", "成都")
	fmt.Printf("切片第二次追加元素后: s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	ss := []string{"武汉", "西安", "苏州"}
	// ... 表示拆分，用于拆分 ss 这个数组，将数组中每一个元素取出，然后追加到切片 s1 中
	s1 = append(s1, ss...)
	fmt.Printf("切片第三次追加元素后: s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	fmt.Println()

	// 删除数组中的元素
	// 声明数组 x1
	x1 := [...]int{1, 3, 5}
	// 通过 数组 x1 得到一个切片 s1
	sss1 := x1[:]
	fmt.Println(sss1, len(sss1), cap(sss1)) // [1 3 5] 3 3

	/*
		1. 切片不保存具体的值
		2. 切片对应一个底层数组
		3. 底层数组都是占用一块连续的内存
	*/
	fmt.Printf("删除切片中元素前的第 0 个元素的内存地址: %p\n", &sss1[0])
	// 删除 切片 s1 中索引为 1 的元素
	sss1 = append(sss1[:1], sss1[2:]...)
	fmt.Printf("删除切片中元素后的第 0 个元素的内存地址: %p\n", &sss1[0])
	fmt.Println(sss1, len(sss1), cap(sss1)) // [1 5] 2 3

	fmt.Println(x1)	// [1 5 5]

	// 修改底层数组
	sss1[0] = 100
	fmt.Println("修改底层数组后的s1:", sss1)	// [100 5]

	fmt.Println()

	// append 删除切片中的某个元素讲解
	m1 :=[...]int{1,3,5,7,9,11,13,15,17}
	k1 := m1[:]

	// 删掉索引为 2 的元素 5
	k1 = append(k1[:2],k1[3:]...)
	fmt.Println(k1)	// [1 3 7 9 11 13 15 17]
	fmt.Println(m1)	// [1 3 7 9 11 13 15 17 17]

	/*
		这里 append 方法 删除掉了 值为 5 的元素，那么底层数组 m1 就相当于
		从元素 7 开始 到最后一个元素 17，整体向前移动一个位置，也就是 7 移动
		到 原来的元素 5 的位置，9移动到原来元素 7 的位置，以此类推。那么到最后
		的元素 17向前移动到 原来的元素 15的位置上以后，原先的 元素 17的这个位置
		上他还是存在一个17，并没有发生改变。
	*/
}
