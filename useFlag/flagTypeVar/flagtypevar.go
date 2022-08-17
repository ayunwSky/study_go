package flagtypevar

import (
	"flag"
	"fmt"
)

// flag.TypeVar
func FlagTypeVar() {
	var (
		name    string
		age     int
		married bool
		weight  float64
	)

	// 接收字符串
	flag.StringVar(&name, "name", "默认名字", "用户名")
	// 接收整型
	flag.IntVar(&age, "age", 18, "用户年龄")
	// 接收布尔类型
	flag.BoolVar(&married, "married", false, "是否已婚？")
	// 接收浮点型
	flag.Float64Var(&weight, "w", 60.0, "体重")
	// 解析
	flag.Parse()
	fmt.Printf("姓名: %v\n", name)
	fmt.Printf("年龄: %v\n", age)
	fmt.Printf("是否已婚？: %v\n", married)
	fmt.Printf("体重: %v\n", weight)
}
