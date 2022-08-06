/**
 * @Author: Allen_Jol
 * @Date: 2022/8/6 14:58
 * @Desc:
 */

package main

import "fmt"

// 接口的实现
/*
接口保存的数据分为两部分：值的动态类型和动态值。
一个接口类型的变量，什么值都可以存
接口是一个引用类型，存放的是动态类型和动态值

比如：该例子中两个类型就是 main.dog 和 main.chicken
值就是我们再声明变量 jinmao 和 kfc 的时候赋的值
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

func (d dog) move() {
	fmt.Println("小狗在跑步")
}

func (d dog) eat(food string) {
	fmt.Printf("小狗在吃%s\n", food)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("小鸡在走路")
}

func (c chicken) eat(food string) {
	fmt.Printf("小鸡在吃%s\n", food)
}

func main() {
	// 定义一个接口类型的变量
	var a1 animal
	fmt.Println(a1)	// nil

	// 定义一个 dog 类型的变量 jinmao
	jinmao := dog{
		name: "金毛",
		feet: 4,
	}

	a1 = jinmao
	fmt.Println(a1)
	fmt.Printf("%T\n", a1)

	// 调用 move 方法和 eat 方法
	a1.move()
	a1.eat("肉骨头")

	fmt.Println("------------------------")

	var a2 animal
	kfc := chicken{
		feet: 2,
	}

	a2 = kfc
	fmt.Println(a2)
	fmt.Printf("%T\n", a2)

	a2.move()
	a2.eat("玉米粒")

}
