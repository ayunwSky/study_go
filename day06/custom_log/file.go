/**
 * @Author: ayunwSky
 * @Date: 2022/8/13 20:10
 * @Desc:
 */

package custom_log

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文文件里面写日志

type FileLogger struct {
	Level      LogLevel
	filePath   string // 日志文件保存的路径
	fileName   string // 日志文件保存的目录
	fileObj    *os.File
	errFileObj *os.File
	//errFileName string // error 级别的日志单独记录一个文件
	maxFileSize int64 // 日志文件的最大大小
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, filePath, fileName string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}

	fl := &FileLogger{
		Level:       logLevel,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: maxSize,
	}
	// 按照文件路径和文件名 将文件打开
	if err := fl.initFile(); err != nil {
		panic(err)
	}

	return fl
}

func (f *FileLogger) initFile() (err error) {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err: %v\n", err)
		return err
	}

	errFileObj, err := os.OpenFile(fullFileName+"_error", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err: %v\n", err)
		return err
	}
	// 日志文件都已经打开
	f.fileObj = fileObj
	f.errFileObj = errFileObj

	return nil
}

func (f *FileLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= f.Level
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getLogInfo(3)
		tf := now.Format("2006-01-02 15:04:05")
		fmt.Fprintf(f.fileObj, "[%s] [%s] [fileName:%s - funcName:%s - lineNo:%d]: %s\n", tf, getLogString(lv), fileName, funcName, lineNo, msg)

		// 如果要记录的日志大于等于ERROR级别，则还要在ERROR日志文件中再记录一遍
		if lv >= ERROR{
		fmt.Fprintf(f.errFileObj, "[%s] [%s] [fileName:%s - funcName:%s - lineNo:%d]: %s\n", tf, getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}

func (f *FileLogger) Debug(format string, a ...interface{}) {
		f.log(DEBUG, format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{}) {
		f.log(INFO, format, a...)
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
		f.log(WARNING, format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{}) {
		f.log(ERROR, format, a...)
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
		f.log(FATAL, format, a...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
