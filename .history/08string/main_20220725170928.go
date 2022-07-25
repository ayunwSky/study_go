package main

import (
	"fmt"
	"strings"
)

/*
Go中字符串必须用双引号包裹
单引号包裹的表示的是字符

字符串：
	s := "Hello 小姜"
字符：
	c1 := 'h'
	c2 := '1'
	c3 := '姜'
字节：1字节=8bit（ 8个二进制位 ）
1个字符'A'=1个字节
1个utf-8编码的汉字 '姜' = 一般占3个字节
*/

func main() {
	// 字符串转义，打印一个windows路径,加 \ 将原来的路径中的 \ 进行转义
	// 双引号中包裹单引号不需要转义
	path1 := "E:\\gopath\\src\\github.com\\allenjol\\08string"
	path2 := "\"E:\\gopath\\src\\github.com\\allenjol\\08string\""
	path3 := "'E:\\gopath\\src\\github.com\\allenjol\\08string'"

	fmt.Println("path1=", path1)
	fmt.Println("path2=", path2)
	fmt.Println("path3=", path3)
	// fmt.Printf("%#v", path)

	s := "I'm ok"
	fmt.Println(s)


	// 多行字符串
	s2 := `
		床前明月光
	床下两双鞋
			君问何处去
	`
	fmt.Println(s2)

	s3 := `原样输出`
	fmt.Println(s3)

	s4 := `'E:\gopath\src\github.com\allenjol\08string'`
	fmt.Println(s4)

	// 字符串常用操作

	// 计算长度
	fmt.Println("len of s3:", len(s3))

	// 字符串拼接
	name := "小姜"
	world := "is boy"
	ss := name + world

	// 字符串拼接方法一
	fmt.Println(name + world)
	fmt.Println(ss)

	// 字符串拼接方法二
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println("ss1:", ss1)

	// 字符串分割
	ret := strings.Split(s4, "\\")
	fmt.Println(ret)

	// 判断是否 包含
	fmt.Println("是否包含:")
	fmt.Println(strings.Contains(ss, "小姜"))
	fmt.Println(strings.Contains(ss, "小王"))

	// 前缀 和 后缀
	fmt.Println("前缀和后缀:")
	fmt.Println(strings.HasPrefix(ss, "小"))
	fmt.Println(strings.HasSuffix(ss,"boy"))

	// 判断 索引位置，第一个值索引为 0
	s5 := "abcede"
	fmt.Println(strings.Index(s5, "c"))
	fmt.Println(strings.LastIndex(s5, "e"))

	// join 拼接
	fmt.Println(strings.Join(ret, "-"))

	fmt.Println()
	// Go中实现字符串修改，默认字符串是不能修改的变量
	s6 := "白萝卜"	// '白' '萝' '卜'
	// 把字符串强制转换成一个rune切片，切片中保存的就是字符
	runeS2 := []rune(s6)
	runeS2[0] = '红'	// 要用 单引号 表示字符
	// 把rune切片强制转换成字符串
	fmt.Println(string(runeS2))

	c1 := "红"	// string
	c2 := '红'	// int32
	fmt.Printf("c1:%T c2: %T\n", c1, c2)

	c3 := "H"	// string
	c4 := 'H'	// int32
	fmt.Printf("c3:%T c4: %T\n", c3, c4)
	fmt.Printf("%")

}
