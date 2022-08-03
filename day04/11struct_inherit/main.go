package main

import "fmt"

// 模拟实现其他语言中的 “继承”

// 动物结构体
type animal struct {
	name string
}

// 狗的结构体
type dog struct {
	feet uint8
	animal	// animal 拥有的方法，dog此时都拥有。变相的模拟实现了一个继承
}

// 给 animal 实现一个动的方法
func (a animal) move() {
	fmt.Printf("%s: 会动\n", a.name)
}

// 给 dog 实现一个叫的方法
func (d dog) speak() {
	fmt.Printf("%s在叫: 汪汪汪\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{
			"金毛",
		},
		feet: 4,
	}
	fmt.Println(d1)
	
	fmt.Println()

	d1.speak()
	d1.move()
}
