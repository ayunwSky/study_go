package main

import "fmt"

type dog struct {
	name string
}

// 构造结构体函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
/*
	接收者 通常写在函数的前面，比如这里的 (d dog)，
	d 表示形参，dog 表示 dog 类型的接收者，一般形参
	都用接收者的首字母的小写，比如这里的 d 就是 dog 类型
	首字母的小写。

	即：接收者表示的是调用该方法的具体类型变量，一般用类型名称的首字母小写来表示

*/

func (d dog) speak() {
	fmt.Printf("%s:汪汪汪~~~", d.name)
}

func main() {
	d1 := newDog("泰迪")
	d1.speak()
}
