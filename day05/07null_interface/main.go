/**
 * @Author: Allen_Jol
 * @Date: 2022/8/7 17:07
 * @Desc:
 */

package main

import "fmt"

/*
空接口
空接口无需起名字，通常定义为：interface{}
空接口是指没有定义任何方法的接口。因此任意类型都实现了空接口
空接口类型的变量可以存储任意类型的变量
*/

/*
interface:关键字
interface{}:空接口类型
当我们不太确定需要用到什么类型的时候，就可以使用空接口

类型断言：
空接口可以存储任意类型的值，那怎么获取其存储的具体数据呢？就通过类型断言
*/

// 2. 空接口接收任意类型的函数参数例子
func show(a interface{}){
	fmt.Printf("type: %T, value: %v\n",a,a)
}

func main() {
	// 1. 空接口作为 map 的值例子
	var studentInfo = make(map[string]interface{}, 16)

	studentInfo["name"] = "allen"
	studentInfo["age"] = 18
	studentInfo["married"] = true
	studentInfo["hobbies"] = [...]string{"唱歌", "开车", "打篮球"}
	fmt.Println(studentInfo)

	fmt.Println()

	show(10)
	show(false)
	show(nil)
	show(studentInfo)
}
