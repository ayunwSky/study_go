/**
 * @Author: ayunwSky
 * @Date: 2022/8/13 20:10
 * @Desc:
 */

package custom_log

import (
	"fmt"
	"time"
)

// 往文文件里面写日志

type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件保存的目录
	errFileName string // error 级别的日志单独记录一个文件
	maxFileSize int64  // 日志文件的最大大小
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, filePath, fileName string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}

	return &FileLogger{
		Level:       logLevel,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: maxSize,
	}
}

func (f *FileLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= f.Level
}

func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getLogInfo(3)
	tf := now.Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] [fileName:%s - funcName:%s - lineNo:%d]: %s\n", tf, getLogString(lv), fileName, funcName, lineNo, msg)
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
	// 往指定的地方写日志
	//if l.Level > DEBUG {
	//	now := time.Now()
	//	tf := now.Format("2006-01-02 15:04:05")
	//	fmt.Printf("[%s] [%s] %s\n", tf, l.Level, msg)
	//}

	if f.enable(DEBUG) {
		log(DEBUG, format, a...)
	}
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	// 往指定的地方写日志
	if f.enable(INFO) {
		//now := time.Now()
		//tf := now.Format("2006-01-02 15:04:05")
		//fmt.Printf("[%s] [INFO] %s\n", tf, msg)
		log(INFO, format, a...)
	}
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	// 往指定的地方写日志
	if f.enable(WARNING) {
		//now := time.Now()
		//tf := now.Format("2006-01-02 15:04:05")
		//fmt.Printf("[%s] [WARNING] %s\n", tf, msg)
		log(WARNING, format, a...)
	}
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	// 往指定的地方写日志
	if f.enable(ERROR) {
		//now := time.Now()
		//tf := now.Format("2006-01-02 15:04:05")
		//fmt.Printf("[%s] [ERROR] %s\n", tf, msg)
		log(ERROR, format, a...)
	}
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	// 往指定的地方写日志
	if f.enable(FATAL) {
		//now := time.Now()
		//tf := now.Format("2006-01-02 15:04:05")
		//fmt.Printf("[%s] [FATAL] %s\n", tf, msg)
		log(FATAL, format, a...)
	}
}
