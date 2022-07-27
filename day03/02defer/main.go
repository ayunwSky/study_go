package main

import "fmt"

/*
	defer 多用于函数结束之前释放资源，比如：关闭文件、关闭数据库连接等
	如果有多个defer语句，则从最后一个defer开始执行，然后执行倒数第二个defer,以此类推
*/

func deferDemo() {
	fmt.Println("start...")
	// defer 把它后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("I'm defer 1...")	// 第三次会执行这个 defer
	defer fmt.Println("I'm defer 2...")	// 第二次会执行这个defer
	defer fmt.Println("I'm defer 3...")	// 第一次会执行这个 defer
	fmt.Println("two...")
	fmt.Println("end...")
}

func main() {
	deferDemo()
}
