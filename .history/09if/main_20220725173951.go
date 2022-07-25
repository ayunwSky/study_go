package main

import "fmt"

func main() {
	age := 19
	if age > 18 {
		fmt.Println("成年人,每天晚上交作业")
	} else {
		fmt.Println("未成年,白天写作业")
	}

	// 多分支判断
	if age > 35 {
		fmt.Println("中年人，保温杯里泡枸杞")
	} else if age > 25 {
		fmt.Println("青年")
	} else if age > 18 {
		fmt.Println("成年人了，要负责人")
	} else {
		fmt.Println("未成年")
	}
}
