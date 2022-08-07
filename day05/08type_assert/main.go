/**
 * @Author: Allen_Jol
 * @Date: 2022/8/7 17:31
 * @Desc:
 */

package main

import "fmt"

/*
为什么要有类型断言？
	类型断言可以让我们知道空接口接收的值是什么

类型断言判断空接口中的值格式：x.(T)
	x：表示类型为 interface{} 的变量
	T：表示断言 x 可能是的类型
该语法返回两个参数：
	第一个值是x转化为T类型后的变量
	第二个值是一个布尔值
如果为true，则表示断言成功，为false则表示断言失败
*/

// 类型断言1
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	// 以下这样这样会出现 panic 报错
	//str:=a.(string)
	//fmt.Println(str)

	// 可以避免出现 panic 的方式
	str, ok := a.(string)
	if !ok {
		fmt.Println("断言失败")
	}else {
		fmt.Println("传递进来的是一个字符串:", str)
	}
}

// 类型断言2
func assign2(a interface{}){
	// 断言多次，需要用判断
	switch v := a.(type) {
	case string:
		fmt.Printf("a是string类型,a的值为: %v\n",v)
	case int:
		fmt.Printf("a是int类型,a的值为: %v\n",v)
	case bool:
		fmt.Printf("a是bool类型,a的值为: %v\n",v)
	default:
		fmt.Printf("unsupport type!")
	}
}

func main() {
	assign(100)
	assign("allen")
	assign(true)

	fmt.Println()

	assign2(100)
	assign2("allen")
	assign2(true)
}
