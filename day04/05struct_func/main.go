package main

import "fmt"

// 构造函数

type person struct {
	name string
	age  int
}

type dog struct {
	name, speak string
	age         int
}

// 构造函数
/*
	首先，结构体是值拷贝，也就是 拷贝了一份副本，结构体数据量大的时候会消耗更多内存；

	构造函数是要返回的是结构体还是结构体指针，一般是这样考虑的:
		1. 当结构体中的字段只有几个的时候，可以直接返回整个结构体
		2. 当结构体中的字段比较多的时候，尽量返回结构体的指针类型，也就是 *person 这种内存地址
			这样可以减少程序的内存开销
	
	所以，个人认为: 还是一开始学习的时候就直接用指针的方法就可以了

*/

// func newPerson(name string, age int) person {
// 	return person{
// 		name: name,
// 		age:  age,
// 	}
// }

// func main() {
// 	p1 := newPerson("小姜", 24)
// 	p2 := newPerson("朱迪", 26)
// 	fmt.Println(p1, p2)
// }

func newDogPoint(name, speak string, age int) *dog {
	return &dog{
		name:  name,
		speak: speak,
		age:   age,
	}
}

func main() {
	d1 := newDogPoint("泰迪", "汪汪", 5)
	d2 := newDogPoint("金毛", "汪汪汪", 6)

	fmt.Println(*d1, *d2)
	fmt.Println(d1.name)
}
