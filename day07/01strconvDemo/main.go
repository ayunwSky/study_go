/**
 * @Author: ayunwSky
 * @Date: 2022/8/14 16:10
 * @Desc:
 */

package main

import (
	"fmt"
	"strconv"
)

func main()  {
	str := "10000"
	// 将 str 转换成 十进制的 int64
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Printf("parseint failed, err:%v\n",err)
		return
	}
	fmt.Printf("%#v %T\n", ret1, ret1)

	// Atoi: 字符串转换为整型
	retInt, _ := strconv.Atoi(str)
	fmt.Printf("%#v %T\n", retInt, retInt)

	// 把数字转换成字符串类型
	i := 97
	ret2 := fmt.Sprintf("%d", i)	// "97"
	fmt.Printf("%#v\n", ret2)

	retStr := strconv.Itoa(i)
	fmt.Printf("%#v %T\n", retStr, retStr)

	// 从字符串中解析出布尔值
	boolStr := "true"
	boolValue, _ := strconv.ParseBool(boolStr)
	fmt.Printf("%#v %T\n", boolValue, boolValue)

	// 从字符串解析出浮点数
	floatStr :="3.1415926"
	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Printf("%#v %T\n", floatValue, floatValue)
}
