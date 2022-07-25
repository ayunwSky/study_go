package main

import "fmt"

func main() {
	var n = 100
	
	// 查看类型
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)

	var s = "hello 小姜"
	fmt.Printf("字符串%s\n", s)
	fmt.Printf("字符串%v\n", s)
	fmt.Printf("字符串%#v\n", s)
}
