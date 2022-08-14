/**
 * @Author: Allen_Jol
 * @Date: 2022/8/13 12:56
 * @Desc:
 */

package main

import (
	"github.com/ayunwSky/study_go/day06/mylogger"
)

// 声明一个全局的接口变量
var log mylogger.Logger

// 测试自己定义的日志库
func main() {
	// 1. 往终端打印日志
	//id := 1101
	//name := "ayunwSky"
	////log := mylogger.NewConsoleLog("INFO")
	//log = mylogger.NewConsoleLog("INFO")	// 终端日志实例
	//
	//for {
	//	log.Debug("这是一条 Debug 级别的测试日志")
	//	log.Info("这是一条 Info 级别的测试日志")
	//	log.Warning("这是一条 Warning 级别的测试日志，id: %d, name: %s", id, name)
	//	log.Error("这是一条 Error 级别的测试日志，id: %d, name: %s", id, name)
	//	log.Fatal("这是一条 Fatal 级别的测试日志\n")
	//
	//	time.Sleep(3 * time.Second)
	//}

	// 2. 日志写入到文件。实现一个日志文件500MB
	//log := mylogger.NewFileLogger("INFO", "./", "access.log", 500*1024*1024)

	// 文件日志实例
	log = mylogger.NewFileLogger("INFO", "./", "access.log", 500*1024*1024)

	for {
		log.Debug("这是一条 Debug 级别的测试日志")
		log.Info("这是一条 Info 级别的测试日志")
		log.Warning("这是一条 Warning 级别的测试日志")
		log.Error("这是一条 Error 级别的测试日志")
		log.Fatal("这是一条 Fatal 级别的测试日志")

		//time.Sleep(3 * time.Second)
	}
}
