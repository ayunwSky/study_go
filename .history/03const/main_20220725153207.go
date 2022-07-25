package main

import "fmt"

// 2、常量
// const pi = 3.1415926
// const e = 2.7182

// 批量声明常量
// const (
// 	pi = 3.1415926
// 	e  = 2.7182
// )

// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同
// const (
// 	n1 = 100
// 	n2	// 100
// 	n3	// 100
// )


// iota
/*
1、iota是go语言的常量计数器，只能在常量的表达式中使用。
2、iota在const关键字出现时将被重置为0
3、const中每新增一行常量声明将使iota计数一次
*/

// const (
// 	a1 = iota	// 0
// 	a2			// 1
// 	a3			// 2
// 	a4			// 3
// )

// 常用的 iota 试题
// 用 _ 跳过某些值，_表示忽略
// const (
// 	b1 = iota	// 0
// 	b2			// 1
// 	_			// 2 被跳过
// 	b3			// 3
// )

// iota声明中插队
// const (
// 	c1 = iota	// 0
// 	c2 = 100	// 100
// 	c3 = iota	// 2
// 	c4			// 3
// )
// const c5 = iota // 0

/*
定义数量级 
这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。
同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。
*/

// const (
// 	_ = iota
// 	KB = 1 << (10 * iota)
// 	MB = 1 << (10 * iota)
// 	GB = 1 << (10 * iota)
// 	TB = 1 << (10 * iota)
// 	EB = 1 << (10 * iota)

// )

// 多个 iota 定义在一行
// const (
// 	x, y = iota + 1, iota +2	// 1, 2
// 	z, d						// 2, 3
// 	m, n						// 3, 4
// )

// const (
// 	d1, d2 = iota + 1, iota +2	// d1=1, d2=2
// 	d3, d4						// d3=2, d4=3
// )

func main()  {
	fmt.Println()
}