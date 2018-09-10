package network

import (
	"gitee.com/nggs/log"
)

var Logger log.ILogger = &log.ConsoleLogger{}

func Set(l log.ILogger) {
	Logger = l
}
