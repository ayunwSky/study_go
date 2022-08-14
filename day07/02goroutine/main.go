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

func hello(i int){
	fmt.Println("hello", i)
}

// 程序启动之后会创建一个 main(主) goroutine 去执行
func main(){
	for i:=0;i<1000;i++{
		go hello(i)
	}
	//go hello(10)	// 开启一个单独的 goroutine 去执行 hello 函数(任务)

	fmt.Println("main")
	// main 函数结束了，则由main函数启动的 goroutine 都结束了
	time.Sleep(time.Second)
}
