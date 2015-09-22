package wlog

import (
	"fmt"
	"log"
	"os"
)

var consoleLogger = log.New(os.Stdout, "", LogFormat)

type ConsoleAppender struct {
	*BaseAppender
}

// NewConsoleAppender
func NewConsoleAppender() *ConsoleAppender {
	m := &ConsoleAppender{}
	m.BaseAppender = NewBaseAppender("console")
	return m
}

// WriteLog implements Appender
// output example: [INFO] mabetle: message.
func (a ConsoleAppender) WriteLog(level string, catalog string, callin int, v ...interface{}) {
	logMsg := fmt.Sprint(v...)
	msg := fmt.Sprintf("\n[%s] %s: %s", level, catalog, logMsg)
	msg = GetLevelColorMsg(level, msg)
	consoleLogger.Output(callin, msg)
}

func ScanAddConsoleAppender(lines []string) {
	if !IsHasAppender("console", lines) {
		return
	}
	AddAppender(NewConsoleAppender(), lines)
}
