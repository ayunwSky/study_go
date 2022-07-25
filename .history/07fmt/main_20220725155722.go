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
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)
}
