package main

import "fmt"

func main() {
	var (
		a = 5
		b = 2
	)

	// 算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a++ // 单独的语句，不能放在等于号右边作为赋值 ===> a = a + 1
	b-- // 单独的语句，不能放在等于号右边作为赋值 ===> b = b - 1

	// 关系运算符
	fmt.Println(a == b) // Go 语言是强类型，相同类型变量才能比较
	fmt.Println(a != b) // 不等于
	fmt.Println(a >= b) // 大于等于
	fmt.Println(a <= b) // 小于等于
	fmt.Println(a > b)  // 大于
	fmt.Println(a < b)  // 小于

	// 逻辑运算
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("要苦逼的上班赚钱活命")
	} else {
		fmt.Println("活着就好了")
	}
	// 如果年龄小于18或者年龄大于60岁
	if age < 18 || age < 60 {
		fmt.Println("要苦逼的上班赚钱活命")
	} else {
		fmt.Println("活着就好了")
	}
	// not 取反，原来为真的就为假，原来为假的就为真
	isMarried := false
	fmt.Println(isMarried)
	fmt.Println(!isMarried)

	// 位运算符(二进制位)
	/*
		5的二进制: 101
		2的二进制:  10
	*/
	// & 按位与：两个二进制位都是1，结果才是1
	fmt.Println(5 & 2) // 0
	// | 按位或：两位有一个是1结果就是1
	fmt.Println(5 | 2) // 7
	// ^ 异或：两位不一样则为1
	fmt.Println(5 ^ 2)
	// << : 将5的二进制数 向左移动1位，101向左移动一位变成 1010
	fmt.Println(5 << 1)  // 10
	fmt.Println(1 << 10) // 1024
	// >> : 将 5的二进制数 向右移动1位, 101向右移动一位变成 10
	fmt.Println(5 >> 1) // 2

	// 1 是int8 类型，最多只能移8位，你移动了10位，变成 10000000000，那结果就变成了0
	// 不会报错，但是结果就是 0
	// m := int8(1)
	// fmt.Println(m << 10)

	// 赋值运算符
	var x int
	x = 10
	x += 1 // 等于 x++
	x -= 1 // 等于 x--
	x *= 1 // 等于 x = x * 1
	x /= 1 // 等于 x = x / 1
	x %= 1 // 等于 x = x % 1
	fmt.Println(x)

	x <<= 2 // x = x << 2
	x &= 2  // x = x & 2
	x |= 3  // x = x | 3
	x ^= 3  // x = x ^ 3
	x >>= 3 // x = x >> 3
	fmt.Println(x)

}
