package print

import (
	"fmt"
	"github.com/dashixiong47/KK_BBS/config"
	"log"
	"runtime"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

var currentLogLevel LogLevel = 0

func SetLogLevel(level LogLevel) {
	currentLogLevel = level
}
func logMessage(level LogLevel, levelName string, format string, a ...interface{}) {
	if level < currentLogLevel {
		return
	}
	mode := config.SettingsConfig.Application.Mode
	if mode == "dev" {
		_, file, line, ok := runtime.Caller(2) // 跳过logMessage函数
		if !ok {
			fmt.Println("error: could not get runtime caller information")
			return
		}

		fmt.Printf("[%s:%d %s] ", file, line, levelName)
		fmt.Printf(format, a...)
		fmt.Println()
	} else {
		log.Print(fmt.Sprintf("[%s] %v", levelName, fmt.Sprintf(format, a...)))
	}
}

func Debug(format string, a ...interface{}) {
	logMessage(DEBUG, "DEBUG", format, a...)
}

func Info(format string, a ...interface{}) {
	logMessage(INFO, "INFO", format, a...)
}

func Warn(format string, a ...interface{}) {
	logMessage(WARN, "WARN", format, a...)
}

func Error(format string, a ...interface{}) {
	logMessage(ERROR, "ERROR", format, a...)
}

// Print 打印日志，并附加文件名、行号和函数名
//func Print(format string, a ...interface{}) {
//	mode := config.SettingsConfig.Application.Mode
//	if mode == "dev" {
//
//		_, file, line, ok := runtime.Caller(1) // 获取调用者的上下文
//		if !ok {
//			fmt.Println("error: could not get runtime caller information")
//			return
//		}
//		//funcName := runtime.FuncForPC(reflect.ValueOf(Print).Pointer()).Name()
//		fmt.Printf("[%s:%d] ", file, line)
//		fmt.Printf(format, a...)
//		fmt.Println()
//	} else {
//		log.Println(fmt.Sprintf(format, a...))
//	}
//
//}
