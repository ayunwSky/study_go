/**
 * @Author: ayunwSky
 * @Date: 2022/8/14 17:14
 * @Desc:
 */

package main

import (
	"fmt"
	"time"
)

// goroutine

/*
	goroutine 的启动是需要等待一定的时间并且消耗资源的
	因此上面的for循环都可能已经执行到不止第一个 第二个的循环，
	goroutine 才开始工作。

	goroutine 什么时候结束？
		goroutine 对应的函数执行结束后，goroutine就结束了
		main 函数执行完后，由 main 函数创建的 goroutine 也就结束了
*/

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后会创建一个 main(主) goroutine 去执行
func main() {
	//for i:=0;i<100;i++{
	//	go hello(i)
	//}
	////go hello(10)	// 开启一个单独的 goroutine 去执行 hello 函数(任务)
	//
	//fmt.Println("main")
	//// main 函数结束了，则由main函数启动的 goroutine 都结束了
	//time.Sleep(time.Second)

	// 使用匿名函数
	//for i:=0;i<100;i++{
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//}
	////go hello(10)	// 开启一个单独的 goroutine 去执行 hello 函数(任务)
	//
	//fmt.Println("main")
	//// main 函数结束了，则由main函数启动的 goroutine 都结束了
	//time.Sleep(time.Second)

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i) // 此时用的 i 是匿名函数内部传递进来的 i
		}(i)
	}
	//go hello(10)	// 开启一个单独的 goroutine 去执行 hello 函数(任务)

	fmt.Println("main")
	// main 函数结束了，则由main函数启动的 goroutine 都结束了
	time.Sleep(time.Second)

}
