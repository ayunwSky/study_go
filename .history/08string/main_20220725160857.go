package main

import "fmt"

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
	path1 := "E:\\gopath\\src\\github.com\\allenjol\\08string"
	fmt.Println("path=", path)
	fmt.Println("path=", path)
	// fmt.Printf("%#v", path)

}
