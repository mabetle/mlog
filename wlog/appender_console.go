package wlog

import (
	"fmt"
	"github.com/mabetle/mcore/mcon"
	"github.com/mabetle/mcore/mterm"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
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

func isWindows() bool {
	goos := strings.ToUpper(runtime.GOOS)
	if strings.Contains(goos, "WINDOW") {
		return true
	}
	return false
}

// WriteLog implements Appender
// output example: [INFO] mabetle: message.
func (a ConsoleAppender) WriteLog(level string, catalog string, callin int, v ...interface{}) {
	if !mterm.IsXterm() && isWindows() {
		// for win cmd
		winCmdOut(level, catalog, callin, v...)
		return
	}
	// xterm and others
	logMsg := fmt.Sprint(v...)
	msg := fmt.Sprintf("\n[%s] %s: %s", level, catalog, logMsg)
	msg = GetLevelColorMsg(level, msg)
	consoleLogger.Output(callin, msg)
}

func winCmdOut(level string, catalog string, callin int, v ...interface{}) {
	level = strings.TrimSpace(level)
	level = strings.ToUpper(level)
	//callin = 3
	msg := fmt.Sprint(v...)
	createTime := time.Now()
	_, file, line, ok := runtime.Caller(callin)
	if !ok {
		file = "???"
		line = 0
	}
	fmt.Fprintf(os.Stderr, "%v %s:%d", createTime, file, line)

	logMsg := fmt.Sprintf("\n[%s] %s: %s\n", level, catalog, msg)

	switch level {
	case "TRACE":
		mcon.PrintMagenta(logMsg)
	case "DEBUG":
		mcon.PrintCyan(logMsg)
	case "INFO":
		mcon.PrintGreen(logMsg)
	case "WARN":
		mcon.PrintYellow(logMsg)
	case "ERROR":
		mcon.PrintRed(logMsg)
	default:
		fmt.Fprint(os.Stderr, logMsg)
	}
}

func ScanAddConsoleAppender(lines []string) {
	if !IsHasAppender("console", lines) {
		return
	}
	AddAppender(NewConsoleAppender(), lines)
}
