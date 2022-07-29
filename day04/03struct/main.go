package main

import "fmt"

// 结构体 struct

/*
	Go语言中的基础数据类型可以表示一些事物的基本属性，
	但是当我们想表达一个事物的全部或部分属性时，这时候再用单一的基本数据类型明显就无法满足需求了
	而 struct 是 Go语言提供了一种自定义数据类型，可以封装多个基本数据类型

	语言内置的基础数据类型是用来描述一个值的，而结构体是用来描述一组值的。
	比如一个人有名字、年龄和居住城市等，本质上是一种聚合型的数据类型

	定义方法：
		type 类型名 struct {
			字段名 字段类型
			字段名 字段类型
			…
		}

	其中：
		类型名：标识自定义结构体的名称，在同一个包内不能重复。
		字段名：表示结构体字段名。结构体中的字段名必须唯一。
		字段类型：表示结构体字段的具体类型。
*/

type ayunw struct {
	name    string
	age     int
	gender  string
	city    string
	hobbies []string
}

func main() {
	// 声明一个 ayunw 类型的变量 jiang
	var jiang ayunw
	// 通过字段赋值（即 结构体实例化）
	jiang.name = "小姜"
	jiang.age = 18
	jiang.gender = "男"
	jiang.city = "浙江"
	jiang.hobbies = []string{"音乐", "篮球", "开车"}

	// 访问变量jiang的字段
	fmt.Printf("%T\n", jiang)
	fmt.Println(jiang)
	fmt.Println(jiang.name)
	fmt.Println(jiang.hobbies)
	fmt.Printf("jiang=%v\n", jiang)
	fmt.Printf("jiang=%#v\n",jiang)

	fmt.Println()

	var wang ayunw
	wang.name = "小王"
	wang.age = 20

	fmt.Printf("%T\n",wang)
	fmt.Println(wang)
	fmt.Println(wang.name)
	fmt.Println(wang.age)

	fmt.Println()

	// 匿名结构体：多用于临时场景，一般就用一次就可能不用了
	// var 声明 s 变量，是一个匿名结构体类型
	var s struct{
		x string
		y int
	}
	s.x = "哈哈"
	s.y = 100

	fmt.Printf("type: %T value:%v\n",s,s)

}
