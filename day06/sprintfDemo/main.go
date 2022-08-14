/**
 * @Author: ayunwSky
 * @Date: 2022/8/14 10:48
 * @Desc:
 */

package main

// 字符串拼接的方法：https://www.cainiaojc.com/golang/go-concatenate-two-strings.html

import (
	"fmt"
	"time"
)

// Sprintf 拼接
func sprintfStringDemo() {
	str1 := "ayunwSky"
	str2 := "Love"
	str3 := "Go"

	result := fmt.Sprintf("%s-%s-%s", str1, str2, str3)
	fmt.Println(result)
}

func sprintfStringIntDemo() {
	//now := time.Now()
	//timeFormat := now.Format("2006-01-02 15:04:05")
	customLogFile := "access.log"
	timeFormat := time.Now().Format("20060102150405")
	fmt.Println(timeFormat)

	result := fmt.Sprintf("%s.%s.bak",customLogFile, timeFormat)
	fmt.Println(result)
}

// Sprintf 拼接不同类型的字符串
func main() {
	sprintfStringDemo()
	fmt.Println()
	sprintfStringIntDemo()
}
