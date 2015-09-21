package mlog

import (
	"github.com/mabetle/mlog/logapi"
	"github.com/mabetle/mlog/wlog"
)

// GetLogger returns logapi.Logger
func GetLogger(catalog string) *wlog.WrapLogger {
	return wlog.NewWrapLogger(catalog, 5)
}

// GetWrapLogger
func GetWrapLogger(catalog string) *wlog.WrapLogger {
	return wlog.NewWrapLogger(catalog, 5)
}

// GetApiLogger
func GetApiLogger(catalog string) logapi.Logger {
	return wlog.NewLogger(catalog, 5)
}
