package mlog

import (
	"github.com/mabetle/mlog/logapi"
	"github.com/mabetle/mlog/wlog"
)

// GetLogger returns wlog.WrapLogger
func GetLogger(catalog string) *wlog.WrapLogger {
	return wlog.NewWrapLogger(catalog, 5)
}

func GetMainLogger() *wlog.WrapLogger {
	return wlog.NewWrapLogger("main", 5)
}

// GetWrapLogger returns wlog.WrapLogger
func GetWrapLogger(catalog string) *wlog.WrapLogger {
	return wlog.NewWrapLogger(catalog, 5)
}

// GetApiLogger valid wlog implements Logger API
func GetApiLogger(catalog string) logapi.Logger {
	return wlog.NewLogger(catalog, 5)
}
