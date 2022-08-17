package flagtype

import (
	"flag"
	"fmt"
)

// flag.Type
func FlagType() {
	var (
		// 接收字符串
		name = flag.String("name", "默认名字", "用户姓名")
		// 接收整形
		age = flag.Int("age", 24, "用户年龄")
		// 接收布尔型
		married = flag.Bool("married", false, "是否已婚")
	)

	// 解析
	flag.Parse()

	fmt.Printf("name: %v\n", *name)
	fmt.Printf("age: %v\n", *age)
	fmt.Printf("married: %v\n", *married)
}
