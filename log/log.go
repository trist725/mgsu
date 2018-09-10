package log

import (
	"fmt"
	l "log"

	"github.com/astaxie/beego/logs"

	"gitee.com/nggs/util"
)

const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

type ILogger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	SetLevel(int)
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	Close()
}

func New(logDir string, logFileBaseName string) ILogger {
	util.MustMkdirIfNotExist(logDir)

	var logger = logs.NewLogger(10000)

	logger.Async()

	//logger.EnableFuncCallDepth(true)

	//logger.SetLogFuncCallDepth(3)

	config := fmt.Sprintf(`{"filename":"%s/%s.log","level":%d,"maxlines":250000,"separate":["error"]}`,
		logDir, logFileBaseName, logs.LevelDebug)

	logger.SetLogger("multifile", config)

	return logger
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type ConsoleLogger struct {
	level int
}

func NewConsoleLogger() (l ILogger) {
	l = &ConsoleLogger{
		level: logs.LevelDebug,
	}
	return
}

func (cl ConsoleLogger) Debug(format string, args ...interface{}) {
	if LevelDebug > cl.level {
		return
	}
	l.Printf("[D] %s\n", fmt.Sprintf(format, args...))
}

func (cl ConsoleLogger) Info(format string, args ...interface{}) {
	if LevelInformational > cl.level {
		return
	}
	l.Printf("[I] %s\n", fmt.Sprintf(format, args...))
}

func (cl ConsoleLogger) Warn(format string, args ...interface{}) {
	if LevelWarning > cl.level {
		return
	}
	l.Printf("[W] %s\n", fmt.Sprintf(format, args...))
}

func (cl ConsoleLogger) Error(format string, args ...interface{}) {
	if LevelError > cl.level {
		return
	}
	l.Printf("[E] %s\n", fmt.Sprintf(format, args...))
}

func (cl *ConsoleLogger) SetLevel(level int) {
	cl.level = level
}

func (ConsoleLogger) Close() {

}
