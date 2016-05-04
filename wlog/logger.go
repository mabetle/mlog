package wlog

import (
	"fmt"

	"github.com/mabetle/mlog/logapi"
)

// WrapLogger .
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

// Info .
func (l WrapLogger) Info(args ...interface{}) { WriteLog("INFO", l.Catalog, l.Callin, args...) }

// Warn .
func (l WrapLogger) Warn(args ...interface{}) { WriteLog("WARN", l.Catalog, l.Callin, args...) }

// Debug .
func (l WrapLogger) Debug(args ...interface{}) { WriteLog("DEBUG", l.Catalog, l.Callin, args...) }

// Error .
func (l WrapLogger) Error(args ...interface{}) { WriteLog("ERROR", l.Catalog, l.Callin, args...) }

// Trace .
func (l WrapLogger) Trace(args ...interface{}) { WriteLog("TRACE", l.Catalog, l.Callin, args...) }

// Infof .
func (l WrapLogger) Infof(format string, args ...interface{}) {
	WriteLog("INFO", l.Catalog, l.Callin, fmt.Sprintf(format, args...))
}

// Warnf .
func (l WrapLogger) Warnf(format string, args ...interface{}) {
	WriteLog("WARN", l.Catalog, l.Callin, fmt.Sprintf(format, args...))
}

// Debugf .
func (l WrapLogger) Debugf(format string, args ...interface{}) {
	WriteLog("DEBUG", l.Catalog, l.Callin, fmt.Sprintf(format, args...))
}

// Errorf .
func (l WrapLogger) Errorf(format string, args ...interface{}) {
	WriteLog("ERROR", l.Catalog, l.Callin, fmt.Sprintf(format, args...))
}

// Tracef .
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
func (l WrapLogger) CheckError(err error, msg ...interface{}) bool {
	okMsg := fmt.Sprint(msg...)

	if err != nil {
		errMsg := fmt.Sprintf("%s Error: %v", okMsg, err)
		WriteLog("WARN", l.Catalog, l.Callin, errMsg)
		return true
	}

	// no error
	if len(msg) != 0 {
		WriteLog("INFO", l.Catalog, l.Callin, okMsg)
	}

	return false
}

// CheckErrorf check errors
func (l WrapLogger) CheckErrorf(err error, msg string, args ...interface{}) bool {
	okMsg := fmt.Sprintf(msg, args...)
	if okMsg != "" {
		return l.CheckError(err, okMsg)
	}
	return l.CheckError(err)
}

// LoadConfig .
func (l WrapLogger) LoadConfig(location string) {
	LoadConfig(location)
}

// SetLevel .
func (l WrapLogger) SetLevel(level string, catalog ...string) {
	SetLevel(level, catalog...)
}
