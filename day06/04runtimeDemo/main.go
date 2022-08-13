/**
 * @Author: Allen_Jol
 * @Date: 2022/8/13 14:05
 * @Desc:
 */

package main

import (
	"fmt"
	"path"
	"runtime"
)

//func main() {
//	// main 函数直接调用runtime，所以参数设置为0即可
//	pc, file, line, ok := runtime.Caller(0)
//	if !ok{
//		fmt.Printf("runtime.Caller() failed\n")
//		return
//	}
//
//	fmt.Println(pc)
//	fmt.Println(file)	// /Users/allenjol/gopath/src/github.com/ayunwSky/study_go/day06/04runtimeDemo/main.go
//	fmt.Println(line)	// 15
//}

//func f1() {
//	// f1 函数调用runtime，所以应该向上找一层，则参数应该设置为 1，此时 f1调用的行数就会被打印
//	pc, file, line, ok := runtime.Caller(1)
//	if !ok{
//		fmt.Printf("runtime.Caller() failed\n")
//		return
//	}
//
//	fmt.Println(pc)
//	fmt.Println(file)	// /Users/allenjol/gopath/src/github.com/ayunwSky/study_go/day06/04runtimeDemo/main.go
//	fmt.Println(line)	// 41
//}
//
//func main(){
//	f1()
//}

func f() {
	/*
		pc: 函数名
		file: 文件名
		line: 行号
	*/
	// f 函数调用runtime，则参数设置为 1，此时 f1 调用的行数就会被打印.
	// 如果想要拿到再 main() 函数中的 f1() 调用的行号，则参数应该设置为 2，以此类推
	// 那么因为是在 main 函数中调用了 f1函数，这里设置为 2 的话，函数名就是 main。
	// 如果想要拿到 函数名 f1，则应该设置为 1
	pc, file, line, ok := runtime.Caller(2) // 设置为1,则 函数名为：main.f1
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	//fmt.Println(file)	// /Users/allenjol/gopath/src/github.com/ayunwSky/study_go/day06/04runtimeDemo/main.go
	fmt.Println(path.Base(file)) // path.Base 传递一个路径，他可以只拿到最后一个文件名
	fmt.Println(line)            // 62
}

func f1() {
	f()
}
func main() {
	f1()
}
