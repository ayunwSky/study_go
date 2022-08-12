/**
 * @Author: Allen_Jol
 * @Date: 2022/8/11 21:33
 * @Desc:
 */

package main

import (
	"fmt"
	"time"
)

func timeDemo() {
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Date())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
}

// 时间戳。从1970年1月1日(08:00:00 GMT)到当前时间的总毫秒数
func timestampDemo1() {
	now := time.Now()
	// 时间戳
	timestampUnix := now.Unix()
	// 纳秒时间戳
	timestampUnixNano := now.UnixNano()

	fmt.Println("timestampUnix:", timestampUnix)
	fmt.Println("timestampUnixNano:", timestampUnixNano)

}

// 用 time.Unix() 函数将时间戳转换成时间格式
func timeUnixToTimeFormat(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)
	fmt.Println("now time:", timeObj)

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minut := timeObj.Minute()
	second := timeObj.Second()

	fmt.Printf("format time: %d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, second, minut)
}

// 时间操作
func timeOperators() {
	now := time.Now()

	// 当前时间加10分钟
	fmt.Println(now.Add(10 * time.Minute))
	// 当前时间加24小时
	fmt.Println(now.Add(24 * time.Hour))

}

func timeTick() {
	timmer := time.Tick(time.Second)
	for t := range timmer {
		// 1秒钟执行一次
		fmt.Println(t)
	}
}

// 时间格式化
/*
	Go语言中格式化时间模板不是常见的Y-m-d H:M:S，而是用了Go语言诞生的时间
	2006年1月2号15点04分，记忆方法：20061234

	格式化的模板为Go诞生时间2006年1月2号15点04分 Mon Jan
	如果想格式化成12小时的方式，则需要指定 PM
*/
func timeFormat() {
	now := time.Now()

	// 24 小时制
	// 2022-08-11 22:14:47 Thu Aug
	fmt.Println(now.Format("2006-01-02 15:04:05 Mon Jan"))

	// 2022-08-11
	fmt.Println(now.Format("2006-01-02"))
	// 2022/08/11 22:14:47
	fmt.Println(now.Format("2006/01/02 15:04:05"))

	fmt.Println()

	// 12 小时制
	// 2022/08/11 10:14:47 PM Thu Aug
	fmt.Println(now.Format("2006/01/02 03:04:05 PM Mon Jan"))

	// 2022/08/11
	fmt.Println(now.Format("2006/01/02"))
	// 2022/08/11 22:14
	fmt.Println(now.Format("2006/01/02 15:04"))
	// 2022/08/11 10:42 PM
	fmt.Println(now.Format("2006/01/02 03:04 PM"))
	// 22:14 2022/08/11
	fmt.Println(now.Format("15:04 2006/01/02"))

	// 精确到毫秒
	fmt.Println(now.Format("2006/01/02 15:04.000"))
	fmt.Println(now.Format("2006/01/02 15:04.999"))

	// 按照对应的格式解析字符串类型的时间 转换成时间戳
	timeObj, err := time.Parse("2006-01-02", "1996-11-01")
	if err != nil {
		fmt.Printf("parse time failed, err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
}

func main() {
	//timeDemo()
	//timestampDemo1()
	//timeUnixToTimeFormat(1660226020)
	//timeOperators()
	//timeTick()
	timeFormat()
}
