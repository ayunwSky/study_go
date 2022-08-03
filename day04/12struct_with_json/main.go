package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json
/* 
	https://sanyuesha.com/2018/05/07/go-json/
	https://studygolang.com/articles/35184
	
	1. 把 Go 语言中的结构体变量转成 json 格式的字符串，称为序列化
	2. 把 json 格式的字符串转换成 Go语言中能识别的结构体，称为反序列化

	注意：只有 struct 中支持导出的 field 才能被 JSON package 序列化，即首字母大写的 field。
*/



type person struct {
	// 注意：这里是person结构体中定义的，所以在 json 这个包做序列化的时候是会出现空 json 的情况
	// 解决这个空 json 数据的方法就是将 name 和 age 都改成 大写字母开头的，这样 json 包就能识别到

	// 前端可能会告诉你，我前前端处理的时候，都是用小写的，你给我都是大写字母开头不行，怎么解决？
	// 用 json 去解析的时候，用反引号中的小写的 name 来代替 Name,可以写多个，比如 json中用name
	// db 数据库中也用 name，ini格式配置文件中也用 name
	// 注意：在 "json:" 后面不能加空格
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

type Message struct{
	Name string
	Body string
}

func main() {
	p1 := person{
		Name: "小姜",
		Age:  18,
	}

	fmt.Println(p1)

	// 序列化 用 json.Marshal
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err: %v", err)
		return
	}
	fmt.Printf("%#v\n", string(b)) // 打印出来会有一个 反斜杠
	fmt.Printf("%v\n", string(b))  // 没有反斜杠

	// 反序列化
	str := `{"name": "allen","age": 24,"hobbies": ["swim", "read"]}`
	var p2 person
	// 要将指针传递进去，为了能在 Unmarshal 函数内部修改 p2 的值 而不是修改副本
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%#v\n", p2)
	fmt.Printf("%v\n", p2)

	fmt.Println()

	m := Message{
		"allen",
		"Hello",
	}

	b2, err := json.Marshal(m)

	if err != nil{
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("%v", string(b2))
}
