/**
 * @Author: Allen_Jol
 * @Date: 2022/8/13 12:52
 * @Desc:
 */

package custom_log

import (
	"fmt"
	"os"
	"time"
)

// 往终端上打印日志

// Logger 日志结构体
type Logger struct {
	Level LogLevel
}

// NewLog 构造函数
func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		fmt.Println("parse log level failed!")
		os.Exit(1)
	}

	return Logger{
		Level: level,
	}
}

// format 和 a 组合起来是一条完整的 msg 日志信息
func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getLogInfo(3)
	tf := now.Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] [fileName:%s - funcName:%s - lineNo:%d]: %s\n", tf, getLogString(lv), fileName, funcName, lineNo, msg)
}

func (l Logger) enable(LogLevel LogLevel) bool {
	return LogLevel >= l.Level
}

func (l Logger) Debug(format string, a ...interface{}) {
	// 往指定的地方写日志
	//if l.Level > DEBUG {
	//	now := time.Now()
	//	tf := now.Format("2006-01-02 15:04:05")
	//	fmt.Printf("[%s] [%s] %s\n", tf, l.Level, msg)
	//}

	if l.enable(DEBUG) {
		log(DEBUG, format, a...)
	}
}

func (l Logger) Info(format string, a ...interface{}) {
	// 往指定的地方写日志
	if l.enable(INFO) {
		//now := time.Now()
		//tf := now.Format("2006-01-02 15:04:05")
		//fmt.Printf("[%s] [INFO] %s\n", tf, msg)
		log(INFO, format, a...)
	}
}

func (l Logger) Warning(format string, a ...interface{}) {
	// 往指定的地方写日志
	if l.enable(WARNING) {
		//now := time.Now()
		//tf := now.Format("2006-01-02 15:04:05")
		//fmt.Printf("[%s] [WARNING] %s\n", tf, msg)
		log(WARNING, format, a...)
	}
}

func (l Logger) Error(format string, a ...interface{}) {
	// 往指定的地方写日志
	if l.enable(ERROR) {
		//now := time.Now()
		//tf := now.Format("2006-01-02 15:04:05")
		//fmt.Printf("[%s] [ERROR] %s\n", tf, msg)
		log(ERROR, format, a...)
	}
}

func (l Logger) Fatal(format string, a ...interface{}) {
	// 往指定的地方写日志
	if l.enable(FATAL) {
		//now := time.Now()
		//tf := now.Format("2006-01-02 15:04:05")
		//fmt.Printf("[%s] [FATAL] %s\n", tf, msg)
		log(FATAL, format, a...)
	}
}
