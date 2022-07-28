package main

import "fmt"

// panic 和 recover
/*
	recover() 必须搭配 defer 使用
	defer一定要在可能引发 panic 的语句之前定义
*/

func funcA() {
	fmt.Println("a")
}

func funcB() {
	// 假设刚才在defer语句前已经打开了数据库连接
	defer func(){
		err := recover()
		fmt.Println(err)

		fmt.Println("释放数据库连接")
	}()
	panic("出现了严重的错误!!!")	// 程序崩溃退出
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
