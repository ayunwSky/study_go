/**
 * @Author: Allen_Jol
 * @Date: 2022/8/6 16:13
 */

package main

import "fmt"

/*
使用值接受者和指针接受者的区别？
	使用值接受者实现接口，结构体类型和结构体指针类型的变量都能存；
	使用指针接收者实现接口，则只能存结构体为指针类型的变量。

通常大多都使用指针接收者的方式
*/

type animal interface {
	move()
	// 接口中可以不写形参，只需要指定类型即可。也就是这里可以不写 food，只需要写类型string即可
	//eat(food string)
	eat(string)
}

type dog struct {
	name string
	feet int8
}

// 使用值接收者实现了接口的所有方法
//func (d dog) move() {
//	fmt.Println("小狗在跑步")
//}
//
//func (d dog) eat(food string) {
//	fmt.Printf("小狗在吃%s\n", food)
//}

// 使用指针接收者实现接口的所有方法
func (d *dog) move() {
	fmt.Println("小狗在跑步")
}

func (d *dog) eat(food string) {
	fmt.Printf("小狗在吃%s\n", food)
}

func main() {
	var a1 animal

	//c1 := dog{ "tom", 4}	// cat
	//c2 := &dog{"jery", 4}	// *cat
	//
	//a1 = c1
	//fmt.Println(a1)
	//a1 = c2
	//fmt.Println(a1)

	c1 := dog{"tom", 4}
	c2 := &dog{"jery", 4}

	a1 = &c1 // 使用指针接收者，也可以在 c1 变量赋值的时候 设置为 &dog，然后这里还是a1 = c1
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
