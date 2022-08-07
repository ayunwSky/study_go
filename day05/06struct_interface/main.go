/**
 * @Author: Allen_Jol
 * @Date: 2022/8/7 16:48
 * @Desc:
 */

package main

import "fmt"

/*
同一个结构体可以实现多个接口，接口还可以嵌套
同一个结构体 dog 实现了 mover 和 eater 两个接口
*/

// 接口嵌套。animal接口嵌套了mover接口和eater接口
type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type dog struct {
	name string
	feet int8
}

// dog 结构体实现了 mover 接口
func (d *dog) move() {
	fmt.Println("小狗在跑步")
}

// dog 结构体实现了 eat 接口
func (d *dog) eat(food string) {
	fmt.Printf("小狗在吃%s\n", food)
}

func main() {
	var a1 animal
	c1 := dog{"tom", 4}
	c2 := &dog{"jery", 4}

	a1 = &c1 // 使用指针接收者，也可以在 c1 变量赋值的时候 设置为 &dog，然后这里还是a1 = c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
