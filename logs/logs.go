package logs

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"time"
)

var (
	logger *logrus.Logger
)
func init() {
	logger = logrus.New()
	//禁止logrus的输出
	src, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	apiLogPath := "./app/app.log"
	logWriter, err := rotatelogs.New(
		apiLogPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(apiLogPath),       // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.ErrorLevel:logWriter,
		logrus.WarnLevel:logWriter,
		logrus.FatalLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})
	logger.AddHook(lfHook)
}

func Info(format string,args ...interface{}){
	_, file, line, ok := runtime.Caller(1)
	message := ""
	if !ok{
		message = "[get file and line err]"
	}else{
		message = fmt.Sprintf("[%s:%v]",file,line)
	}
	logger.Info(message,fmt.Sprintf(format,args...))
}

func Warn(format string,args ...interface{}){
	_, file, line, ok := runtime.Caller(1)
	message := ""
	if !ok{
		message = "[get file and line err]"
	}else{
		message = fmt.Sprintf("[%s:%v]",file,line)
	}
	logger.Warn(message,fmt.Sprintf(format,args...))
}

func Error(format string,args ...interface{}){
	_, file, line, ok := runtime.Caller(1)
	message := ""
	if !ok{
		message = "[get file and line err]"
	}else{
		message = fmt.Sprintf("[%s:%v]",file,line)
	}
	logger.Error(message,fmt.Sprintf(format,args...))
}
