package main

import "fmt"

// 结构体嵌套

// type address struct {
// 	province string
// 	city     string
// }

// type person struct {
// 	name string
// 	age  int
// 	addr address
// 	// province string
// 	// city     string
// }

// type company struct {
// 	name string
// 	addr address
// 	// province string
// 	// city     string
// }

// func main() {
// 	p1 := person{
// 		name: "艾伦",
// 		age:  18,
// 		addr: address{
// 			province: "浙江",
// 			city:     "杭州",
// 		},
// 	}
// 	fmt.Println(p1)
// 	fmt.Println()

// 	fmt.Println(p1.name)
// 	fmt.Println(p1.age)
// 	fmt.Println(p1.addr)
// 	fmt.Println(p1.addr.province)
// 	fmt.Println(p1.addr.city)

// }

// type address struct {
// 	province string
// 	city     string
// }

// type person struct {
// 	name string
// 	age  int
// 	address	// 匿名嵌套结构体 相当于 address: address
// }

// type company struct {
// 	name string
// 	addr address
// }

// func main() {
// 	p1 := person{
// 		name: "艾伦",
// 		age:  18,
// 		address: address{
// 			province: "浙江",
// 			city:     "杭州",
// 		},
// 	}

// 	fmt.Println(p1)
// 	fmt.Println()

// 	fmt.Println(p1.address.city)
// 	fmt.Println(p1.address)
// 	// 先在自己的结构体(person)找city这个字段，找不到再去嵌套的匿名结构体(address)中找city字段
// 	fmt.Println(p1.city)

// }

type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	address
	workPlace
}

type company struct {
	name string
	address
}

func main() {
	p1 := person{
		name: "艾伦",
		age:  18,
		address: address{
			province: "浙江",
			city:     "杭州",
		},
		workPlace: workPlace{
			province: "浙江",
			city:     "衢州",
		},
	}

	fmt.Println(p1)
	fmt.Println()

	fmt.Println(p1.address.city)
	fmt.Println(p1.address)
	// 先在自己的结构体(person)找city这个字段，找不到再去嵌套的匿名结构体(address)中找city字段
	// 因为结构体嵌套了两个匿名结构体 address 和 workPlace,这两个结构体中存在两个相同的字段
	// 所以匿名嵌套结构体出现了嵌套冲突，因此无法使用下面的方法进行打印
	// fmt.Println(p1.city)

	// 防止字段冲突，所以只能用下面的方式来指定打印city是属于哪个结构体中的
	fmt.Println(p1.address.city)
	fmt.Println(p1.workPlace.city)
}
