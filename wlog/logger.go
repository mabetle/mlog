package wlog

import (
	"fmt"
	"github.com/mabetle/mlog/logapi"
)

// WrapLogger
type WrapLogger struct {
	Catalog string
	Callin  int
}

// NewWrapLogger create WrapLogger instanse.
// WrapLogger implemnts api.Logger
func NewWrapLogger(catalog string, callin int) *WrapLogger {
	return &WrapLogger{Catalog: catalog, Callin: callin}
}

// NewLogger returns logapi.Logger
func NewLogger(catalog string, callin int) logapi.Logger {
	return &WrapLogger{Catalog: catalog, Callin: callin}
}

func (l WrapLogger) Info(args ...interface{})  { WriteLog("INFO", l.Catalog, l.Callin, args...) }
func (l WrapLogger) Warn(args ...interface{})  { WriteLog("WARN", l.Catalog, l.Callin, args...) }
func (l WrapLogger) Debug(args ...interface{}) { WriteLog("DEBUG", l.Catalog, l.Callin, args...) }
func (l WrapLogger) Error(args ...interface{}) { WriteLog("ERROR", l.Catalog, l.Callin, args...) }
func (l WrapLogger) Trace(args ...interface{}) { WriteLog("TRACE", l.Catalog, l.Callin, args...) }

func (l WrapLogger) Infof(format string, args ...interface{}) {
	WriteLog("INFO", l.Catalog, l.Callin, fmt.Sprintf(format, args...))
}

func (l WrapLogger) Warnf(format string, args ...interface{}) {
	WriteLog("WARN", l.Catalog, l.Callin, fmt.Sprintf(format, args...))
}

func (l WrapLogger) Debugf(format string, args ...interface{}) {
	WriteLog("DEBUG", l.Catalog, l.Callin, fmt.Sprintf(format, args...))
}

func (l WrapLogger) Errorf(format string, args ...interface{}) {
	WriteLog("ERROR", l.Catalog, l.Callin, fmt.Sprintf(format, args...))
}

func (l WrapLogger) Tracef(format string, args ...interface{}) {
	WriteLog("TRACE", l.Catalog, l.Callin, fmt.Sprintf(format, args...))
}

// CheckNil used in if nil == v statement.
func (l WrapLogger) CheckNil(v interface{}) bool {
	if nil == v {
		msg := "value is nil"
		WriteLog("DEBUG", l.Catalog, l.Callin, msg)
		return true
	}
	return false
}

// CheckError check error and log it.
// if err not nil return true.
func (l WrapLogger) CheckError(err error) bool {
	if err != nil {
		msg := fmt.Sprintf("Error: %v", err)
		WriteLog("WARN", l.Catalog, l.Callin, msg)
		return true
	}
	return false
}

// LoadConfig
func (l WrapLogger) LoadConfig(location string) {
	LoadConfig(location)
}

// SetLevel
func (l WrapLogger) SetLevel(level string, catalog ...string) {
	SetLevel(level, catalog...)
}
