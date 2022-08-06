package main

import "fmt"

// 接口例子2

/*
不管什么牌子的类型的车都能在路上开
*/

// 定义一个car接口类型
// 无论什么结构体，只要有 run 方法，都是属于 car 类型
type car interface {
	run()
}

type ferrari struct {
	brand string
}

type porsche struct {
	brand string
}

func (f ferrari) run() {
	fmt.Printf("%s速度380迈\n", f.brand)
}

func (p porsche) run() {
	fmt.Printf("%s速度350迈\n", p.brand)
}

// drive 函数接收了一个 car 类型变量
// car 是一个接口类型
func drive(c car) {
	c.run()
}

func main() {
	var f1 = ferrari{
		brand: "法拉利",
	}
	var p1 = porsche{
		brand: "保时捷",
	}
	drive(f1)
	drive(p1)
}
