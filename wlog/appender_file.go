package wlog

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var fileLogger = log.New(ioutil.Discard, "", LogFormat)

type FileAppender struct {
	*BaseAppender
	Location string
}

func NewFileAppender(location string) *FileAppender {
	m := &FileAppender{Location: location}
	m.BaseAppender = NewBaseAppender("file")
	return m
}

// WriteLog implements Appender
// output example: [INFO] mabetle: message.
func (a *FileAppender) WriteLog(level string, catalog string, callin int, v ...interface{}) {
	logMsg := fmt.Sprint(v...)
	msg := fmt.Sprintf("\n[%s] %s: %s", level, catalog, logMsg)
	msg = GetLevelColorMsg(level, msg)
	getFileLogger(a).Output(callin, msg)
}

// getFileLogger
func getFileLogger(fa *FileAppender) *log.Logger {
	w, err := os.OpenFile(fa.Location, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Println("Failed to open log file", fa.Location, ":", err)
		return fileLogger
	}
	fileLogger.SetOutput(w)
	return fileLogger
}

// ScanAddFileAppender
func ScanAddFileAppender(lines []string) {
	if !IsHasAppender("file", lines) {
		return
	}
	fileAppendLocation, _ := ScanConfigValue("appender-file-location", lines)
	// add file appender
	if fileAppendLocation != "" {
		//fmt.Printf("fileAppendLocation: %s\n",fileAppendLocation )
		AddAppender(NewFileAppender(fileAppendLocation), lines)
	}
}
