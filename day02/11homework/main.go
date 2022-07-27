package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 1. 判断字符串中汉字的数量,难点在于判断一个字符是汉字
	s1 := "hello浙江挺好"
	count := 0
	// 1. 依次拿到字符串中的字符
	for _, c := range s1 {
		// 2. 判断当前这个字符是不是汉字
		if unicode.Is(unicode.Han, c) {
			// 3. 统计汉字出现的次数
			count++
		}
	}
	fmt.Println(count)

	fmt.Println()

	// 统计单词出现的次数
	s2 := "how do you do"
	// 字符串按空格切割，得到切片
	s3 := strings.Split(s2, " ")
	// 遍历切片，存储到 map
	m1 := make(map[string]int, 10)
	for _, w := range s3 {
		fmt.Println(w)
		// 如果原来 map 中不存在 w 这个key，则出现次数就等于1
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
		// fmt.Println(m1)
		// 累加统计出现次数
		for k, v := range m1 {
			fmt.Println(k, v)
		}
	}

	fmt.Println()

	// 回文
	/*
		上海自来水来自海上
		山西运煤车煤运山西

		思路：
		山	ss[0]  ss[len(ss)-1]
		西	ss[1]  ss[len(ss)-1-1]
		运	ss[2]  ss[len(ss)-1-2]
		煤	ss[3]  ss[len(ss)-1-3]
		车
		煤
		运
		山
		西
	*/

	ss := "上海自来水来自海上"
	// 造一个[]rune 类型的切片
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	fmt.Println(r)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")

}
