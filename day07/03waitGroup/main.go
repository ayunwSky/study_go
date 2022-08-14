/**
 * @Author: ayunwSky
 * @Date: 2022/8/14 17:37
 * @Desc:
 */

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func f() {
	// 设置一个随机种子
	rand.Seed(time.Now().UnixNano())

	for i := 0; i <= 10; i++ {
		r1 := rand.Int()	// int64 的随机数
		r2 := rand.Intn(101) // 设置固定大小的随机数，大于等于0，小于100
		fmt.Println(r1, r2)

		// 设置为 负数
		//fmt.Println(0-r1, 0-r2)
	}
}

func f1(i int){
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

var wg sync.WaitGroup
func main() {
	//f()
	for i:=0;i<10;i++{
		wg.Add(1)	// 计数器登记，计数器加1
		go f1(i)
	}
	// 如何知道这十个goroutine都结束了？
	wg.Wait()	// 等待 wg 的计数器减为 0
}
