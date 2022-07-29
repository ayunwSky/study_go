package main

import "fmt"

// 结构体是值类型
/*
值类型也就是 ctrl+c 和 ctrl+v，生成了一个副本
更改源文件不会影响副本文件的数据
*/

type person struct {
	name, gender string
	age          int
}

// Go 语言中函数传参永远传递的都是一个拷贝出来的副本，更改副本数据不影响源数据
// 如果需要修改 源 数据，那么就可以传递指针过来
func f(x person) {
	x.gender = "女" // 修改的是副本的 gender
	fmt.Println("f 函数中的gender:", x.gender)
}

// 函数 f2 中 传递一个 x 为 person 类型的指针
func f2(x *person) {
	// (*x).gender = "女"	// 根据内存地址找到了 源 变量，这样就修改了源变量的值
	// 和上面的 (*x).gender = "女" 写法是一样的
	x.gender = "女" // 根语法糖，自动根据指针找对应的变量
	fmt.Println("f2 函数中的gender:", x.gender)
}

func main() {
	var p person

	p.name = "艾伦"
	p.gender = "男"
	p.age = 18

	f(p)
	fmt.Println("main 函数中的gender:", p.gender)

	f2(&p) // 把 p 的内存地址传到函数 f2 中
	fmt.Println("main 函数中的gender:", p.gender)

	fmt.Println()

	// 结构体指针 1
	// 用 new 的方法为 struct 初始化，开辟一块内存
	var p2 = new(person)
	(*p2).name = "姜姜"
	p2.gender = "保密" // 直接用语法糖的方式赋值

	fmt.Printf("%T\n", p2)
	fmt.Printf("p2 保存的值的内存地址:%p\n", p2) // p2 保存的值就是一个内存地址
	fmt.Printf("p2 变量的内存地址: %p\n", &p2) // p2 变量的内存地址

	fmt.Println()

	// 结构体指针 2
	// 2.1 key value 类型初始化 (推荐都用这种方式)
	// 这种定义方式中，上面全局结构体定义了 name, age, gender，
	// 而我这里如果只写 name 和 age 是没问题的，少一个字段没关系
	var p3 = person{
		name: "姜总监",
		// gender: "男",
		age: 18,
	}
	fmt.Printf("%#v\n", p3)

	fmt.Println()

	// 2.2 使用值列表的形式初始化
	/*
		注意：
		1. 必须初始化结构体的所有字段，如果缺少某个字段的值会报错；
		2. 初始值的填充顺序必须与字段在结构体中的声明顺序一致
		3. 该方式不能和键值初始化方式混用
	*/
	p4 := person{
		"朱迪",
		"女",
		20,
	}
	fmt.Printf("%#v\n", p4)

	fmt.Println()

	// 取结构体的指针(内存)地址实例化
	p5 := &person{
		name:   "王经理",
		gender: "女",
		age:    26,
	}
	fmt.Println(*p5)
	fmt.Println(p5.name)

	fmt.Println()

	// 结构体占用一块连续的内存空间
	type x struct {
		a int8 // 8bit = 1 byte
		b int8
		c int8
	}

	m := x{
		a: int8(10),
		b: int8(20),
		c: int8(30),
	}
	fmt.Printf("%p\n", &(m.a)) // 0xc0000aa0ca
	fmt.Printf("%p\n", &(m.b)) // 0xc0000aa0cb
	fmt.Printf("%p\n", &(m.c)) // 0xc0000aa0cc
}
