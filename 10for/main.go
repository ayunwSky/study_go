package main

import "fmt"

func main() {
	// Go 语言中循环类型只有for循环
	// for循环的基本格式
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// fmt.Println()

	// 省略初始语句
	// i :=5
	// for ;i<10;i++{
	// 	fmt.Println(i)
	// }

	// 省略初始语句和结束语句
	// i :=5
	// for ; i<10;{
	// 	fmt.Println(i)
	// 	i++
	// }

	// 死循环,不要轻易尝试哦，电脑会死机的！
	// for {
	// 	fmt.Println("123")
	// }

	// for range 键值循环
	/*
		Go语言中可以使用for range遍历数组、切片、字符串、map 及通道（channel）
		通过for range遍历的返回值有以下规律：
			1、数组、切片、字符串返回索引和值。
			2、map返回键和值。
			3、通道（channel）只返回通道内的值。
	*/

	// 一个中文大约占 3 到 4个字节，所以小字索引是6，姜字索引就到了9
	// s := "hello 小姜"
	// for k, v := range s {
	// 	fmt.Printf("%d %c\n", k, v)
	// }

	// 九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, j*i)
		}
		fmt.Println()
	}

}
