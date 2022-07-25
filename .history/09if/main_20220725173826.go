package main

import "fmt"

func main() {
	age := 19
	if age > 18 {
		fmt.Println("成年了")
	} else {
		fmt.Println("未成年")
	}

	// 多分支判断
	if age > 35 {
		fmt.Println("中年人")
	}else if age > 25{
		fmt.Println("青年")
	}
}
