package wlog

import (
	"fmt"
	"log"
	"os"
)

var consoleLogger = log.New(os.Stdout, "", LogFormat)

type ConsoleAppender struct {
}

func NewConsoleAppender() Appender {
	return &ConsoleAppender{}
}

// WriteLog implements Appender
// output example: [INFO] mabetle: message.
func (a ConsoleAppender) WriteLog(level string, catalog string, callin int, v ...interface{}) {
	logMsg := fmt.Sprint(v...)
	msg := fmt.Sprintf("\n[%s] %s: %s", level, catalog, logMsg)
	msg = GetLevelColorMsg(level, msg)
	consoleLogger.Output(callin, msg)
}

// AddConsoleAppender
func AddConsoleAppender() {
	AddAppender("Console", NewConsoleAppender())
}
