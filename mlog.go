package mlog

import (
	"github.com/mabetle/mlog/logapi"
	"github.com/mabetle/mlog/wlog"
)

// SetAppenderLevel
func SetAppenderLevel(appenderName, level string, catalogs ...string) {
	wlog.SetAppenderLevel(appenderName, level, catalogs...)
}

// Set all appenders catalogs level.
func SetLevel(level string, catalogs ...string) { wlog.SetLevel(level, catalogs...) }

// shortcuts for set all appenders level
func SetTraceLevel(catalogs ...string) { SetLevel("trace", catalogs...) }
func SetDebugLevel(catalogs ...string) { SetLevel("debug", catalogs...) }
func SetInfoLevel(catalogs ...string)  { SetLevel("info", catalogs...) }
func SetWarnLevel(catalogs ...string)  { SetLevel("warn", catalogs...) }
func SetErrorLevel(catalogs ...string) { SetLevel("error", catalogs...) }

// AddAppender
func AddAppender(appender logapi.Appender, lines []string) {
	wlog.AddAppender(appender, lines)
}

// LoadConfig
func LoadConfig(location string) {
	wlog.LoadConfig(location)
}
