/**
 * @Author: Allen_Jol
 * @Date: 2022/8/13 12:35
 * @Desc:
 */

package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// 日志例子

/*
	1. 让日志支持往不同的地方输出
	2. 日志级别区分
		1）Debug：调试日志
		2）Trace:
		3）Info: 正常日志
		4）Warning：警告日志
		5）Error： 错误日志
		6）Fatal：严重的错误，程序终止
	3. 日志支持开关控制，即可以支持配置记录哪种级别的日志。如：
		项目开发过程中什么级别都能输出，而项目上线之后，只有INFO及以下级别的日志才能输出
	4. 完整的日志记录要包括：时间、行号、文件名、日志级别、日志信息
	5. 日志文件要切割
*/

func logToStdout() {
	// 1. 直接打印到终端
	for {
		log.Println("这是一条测试日志")
		time.Sleep(3 * time.Second)
	}
}

func logToFile(){
	// 2. 配置日志写到文件中
	fileObj, err := os.OpenFile("./logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	// 设置输出日志到文件，默认输出到 os.stdout 终端
	//log.SetOutput(os.Stdout)
	log.SetOutput(fileObj)

	for {
		log.Println("这是一条测试日志")
		time.Sleep(3 * time.Second)
	}
}

func main() {
	//logToStdout()
	logToFile()
}
