package main

import "fmt"

// panic 和 recover
/*
	panic 对异常进行捕获，类似Python中的 try ... except
	recover() 必须搭配 defer 使用
	defer一定要在可能引发 panic 的语句之前定义

	利用recover可以捕获panic异常，保证程序可以正常执行，不会崩溃导致直接退出程序

	https://www.jianshu.com/p/6c3ce7406009
	https://yuerblog.cc/2019/09/26/golang-%E5%88%A9%E7%94%A8recover%E6%8D%95%E8%8E%B7panic%E5%BC%82%E5%B8%B8/
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
