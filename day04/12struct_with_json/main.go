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

	在 Go 中并不是所有的类型都能进行序列化：

	1) JSON object key 只支持 string
	2) Channel、complex、function 等 type 无法进行序列化
	3) 数据中如果存在循环引用，则不能进行序列化，因为序列化时会进行递归
	4) Pointer 序列化之后是其指向的值或者是 nil

	注意：只有 struct 中支持导出的 field 才能被 JSON package 序列化，即首字母大写的 field。
	但是JSON 对象一般都是小写表示，但是Marshal后寄送对象的首字母仍然还是大写。那我们通常都要用小写的怎么解决？
	使用 struct tags 来解决。

	Struct tag 可以决定 Marshal 和 Unmarshal 函数如何序列化和反序列化数据。
*/

/*
注意：这里是 person 结构体中定义的，所以在 json 这个包做序列化的时候是会出现空 json 的情况
解决这个空 json 数据的方法就是将 name 和 age 都改成 大写字母开头的，这样 json 包就能识别到

前端可能会告诉你，我前端处理的时候，都是用小写的，你给我都是大写字母开头不行，怎么解决？
使用 struct tags，即用 json 去解析的时候，用反引号中的小写的 name 来代替 Name,
可以写多个，比如 json中用name，db 数据库中也用 name，ini格式配置文件中也用 name
注意：在 "json:" 后面不能加空格
*/

//type person struct {
//	Name string `json:"name" db:"name" ini:"name"`
//	Age  int    `json:"age"`
//}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Message struct {
	Name string `json:"name"`
	Body string `json:"body"`
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
	fmt.Println(p2.Name)
	fmt.Println(p2.Age)

	fmt.Println()

	m := Message{
		"allen",
		"Hello",
	}

	b2, err := json.Marshal(m)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("%v", string(b2))
}
