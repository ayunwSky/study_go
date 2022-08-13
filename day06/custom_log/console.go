/**
 * @Author: Allen_Jol
 * @Date: 2022/8/13 12:52
 * @Desc:
 */

package custom_log

import (
	"fmt"
	"time"
)

// 往终端上打印日志

// ConsoleLogger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewLog 构造函数
func NewLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		fmt.Println("parse log level failed!")
		panic(err)
	}

	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= c.Level
}

// format 和 a 组合起来是一条完整的 msg 日志信息
func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getLogInfo(3)
		tf := now.Format("2006-01-02 15:04:05")
		fmt.Printf("[%s] [%s] [fileName:%s - funcName:%s - lineNo:%d]: %s\n", tf, getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	// 往指定的地方写日志
	//if c.Level > DEBUG {
	//	now := time.Now()
	//	tf := now.Format("2006-01-02 15:04:05")
	//	fmt.Printf("[%s] [%s] %s\n", tf, c.Level, msg)
	//}

	//if c.enable(DEBUG) {
	//	log(DEBUG, format, a...)
	//}

	c.log(DEBUG, format, a...)
}

func (c ConsoleLogger) Info(format string, a ...interface{}) {
	// 往指定的地方写日志
	//if c.enable(INFO) {
	//	//now := time.Now()
	//	//tf := now.Format("2006-01-02 15:04:05")
	//	//fmt.Printf("[%s] [INFO] %s\n", tf, msg)
	//	log(INFO, format, a...)
	//}

	c.log(INFO, format, a...)
}

func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	// 往指定的地方写日志
	//if c.enable(WARNING) {
	//	//now := time.Now()
	//	//tf := now.Format("2006-01-02 15:04:05")
	//	//fmt.Printf("[%s] [WARNING] %s\n", tf, msg)
	//	log(WARNING, format, a...)
	//}

	c.log(WARNING, format, a...)
}

func (c ConsoleLogger) Error(format string, a ...interface{}) {
	// 往指定的地方写日志
	//if c.enable(ERROR) {
	//	//now := time.Now()
	//	//tf := now.Format("2006-01-02 15:04:05")
	//	//fmt.Printf("[%s] [ERROR] %s\n", tf, msg)
	//	log(ERROR, format, a...)
	//}

	c.log(ERROR, format, a...)
}

func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	// 往指定的地方写日志
	//if c.enable(FATAL) {
	//	//now := time.Now()
	//	//tf := now.Format("2006-01-02 15:04:05")
	//	//fmt.Printf("[%s] [FATAL] %s\n", tf, msg)
	//	log(FATAL, format, a...)
	//}

	c.log(FATAL, format, a...)
}
