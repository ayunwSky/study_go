package main

import "fmt"

// 引出接口的实例

/*
	接口是一种特殊的类型，它规定了变量有哪些方法。
	编程中存在一种场景：
		我不关心一个变量是什么类型，只关心能调用它的什么方法
*/

type cat struct{}

type dog struct{}

type person struct{}

// 接口是用于规定方法的
// 定义一个能叫的接口类型
type speaker interface {
	// 只要实现了 speak() 方法的变量都是 speaker 类型
	// 这杯称为方法的签名，如果存在多个方法，则这里可以写多个
	speak()
}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪~")
}

func (p person) speak() {
	fmt.Println("啊啊啊~")
}

// 不管传递什么类型的参数，我只关心传递进来的参数可以让我调用 speak 的方法
/*
	假设我这里不定义接口类型而使用 struct 结构体类型，那么我这里只能传递 cat
	或者 dog 或者 person 三种结构体其中的一种。就无法实现我任意传递某一种类型
	都能让 da 这个方法去进行接收并处理
*/
func da(x speaker) {
	// 接收一个参数，传递进来什么就打什么，挨打就要进行 speak
	x.speak()
}

func main() {
	var (
		c1 cat
		d1 dog
		p1 person
	)

	da(c1)
	da(d1)
	da(p1)
}
