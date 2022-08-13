/**
 * @Author: Allen_Jol
 * @Date: 2022/8/13 12:52
 * @Desc:
 */

package custom_log

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
	"time"
)

// 自定义日志

// LogLevel 自定义日志等级的类型别名
type LogLevel uint16

// 定义日志级别
const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		//fmt.Println("日志级别非所设定级别!")
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		//fmt.Println("传入日志级别错误")
		return "INFO"
	}
}

/*
使用 runtime 包中的Caller()函数，来实现 谁调用我，我就返回给他函数名，文件名和行号
*/
func getLogInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]

	return
}

// format 和 a 组合起来是一条完整的 msg 日志信息
func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getLogInfo(3)
	tf := now.Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] [fileName:%s - funcName:%s - lineNo:%d]: %s\n", tf, getLogString(lv), fileName, funcName, lineNo, msg)
}