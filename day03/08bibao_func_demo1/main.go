package main

import (
	"fmt"
	"strings"
)

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		// 判断传递过来的 name 如果不是以 suffix 为后缀的，则给它加上这个 suffix 后缀
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		// 如果 name 是 以 suffix 为后缀的，则 return name
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")

	fmt.Println(jpgFunc("test")) // test.jpg
	fmt.Println(txtFunc("test")) // test.txt

	fmt.Println()
	
	fmt.Println(jpgFunc("haha.jpg")) // haha.jpg
	fmt.Println(txtFunc("ff.txt"))   // ff.txt
}
