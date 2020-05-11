package log

import mlog "github.com/trist725/mgsu/log"

var (
	logger = mlog.NewConsoleLogger()
)

func Logger() mlog.ILogger {
	return logger
}

func SetLogger(l mlog.ILogger) {
	logger = l
}
