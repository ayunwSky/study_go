package main

import "fmt"

// 将 int 定义为自定义 myInt 类型加方法
// 不能给别的包里的类型添加方法，只能给自己的包里的类型添加自定义方法
type myInt int

// 为 myInt 添加一个 hello方法
func (m myInt) hello() {
	fmt.Println("我是自定义类型的 int...")
}

func main() {
	m := myInt(100)
	m.hello()
}
