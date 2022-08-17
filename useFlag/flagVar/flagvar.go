package flagvar

import (
	"flag"
	"fmt"
	"strings"
)

// 自定义变量，并实现flag.Value接口
type Likes []string

func (l *Likes) String() string {
	return fmt.Sprintf("%v", *l)
}

func (l *Likes) Set(s string) error {
	// 分割字符串
	split := strings.Split(s, ",")
	*l = split
	return nil
}

// flag.Var
// 通过 flag.Var() 绑定自定义类型，自定义类型需要实现 Value 接口(Receives必须为指针)
func FlagVar() {
	var likeList Likes
	// 接收自定义类型
	flag.Var(&likeList, "likes", "接收自定义类型")
	// 解析
	flag.Parse()
	fmt.Println(likeList)
}
