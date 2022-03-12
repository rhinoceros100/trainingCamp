package log

import (
	"fmt"
	logs "github.com/astaxie/beego/tree/develop/adapter/logs"
	"strings"
)

var Logger *logs.BeeLogger

func Debug(v ...interface{}) {
	Logger.Debug(generateFmtStr(len(v)), v...)
}

func Warn(v ...interface{}) {
	Logger.Warn(generateFmtStr(len(v)), v...)
}

func Trace(v ...interface{}) {
	Logger.Trace(generateFmtStr(len(v)), v...)
}

func Info(v ...interface{}) {
	Logger.Info(generateFmtStr(len(v)), v...)
}

func Error(v ...interface{}) {
	Logger.Error(generateFmtStr(len(v)), v...)
}

func generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}

func Init(adapterName string, fileName string) error {
	Logger = logs.NewLogger(100000)
	Logger.EnableFuncCallDepth(true)
	Logger.SetLogFuncCallDepth(3)
	if adapterName == "console" {
		return Logger.SetLogger(adapterName, "")
	}
	config := fmt.Sprintf(`{"filename" : "%s"}`, fileName)

	//log.SetLogger(adapterName, `{"filename":"../log/log.txt"}`)
	Logger.DelLogger("console")
	return Logger.SetLogger(adapterName, config)
}
