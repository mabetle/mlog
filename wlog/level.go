package wlog

import (
	"strings"
)

type Level int

const (
	LevelUknown Level = iota
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)

func GetLevelLabel(level Level) string {
	label := "????"
	switch level {
	case LevelTrace:
		label = "TRACE"
	case LevelDebug:
		label = "DEBUG"
	case LevelInfo:
		label = "INFO"
	case LevelWarn:
		label = "WARN"
	case LevelError:
		label = "ERROR"
	default:
	}
	return label
}

func GetStringLevel(str string) Level {
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	level := LevelInfo
	switch str {
	case "trace":
		level = LevelTrace
	case "debug":
		level = LevelDebug
	case "info":
		level = LevelInfo
	case "warn":
		level = LevelWarn
	case "error":
		level = LevelError
	default:
	}
	return level
}

// GetLabelLevel same as GetStringLevel
func GetLabelLevel(label string) Level {
	return GetStringLevel(label)
}

// SetLevel implements api
func SetLevel(level string, catalogs ...string) {
	for _, a := range GetAppenders() {
		a.SetLevel(level, catalogs...)
	}
}

// SetAppenderLevel
func SetAppenderLevel(appenderName, level string, catalogs ...string) {
	appenderName = strings.ToLower(appenderName)
	for n, a := range GetAppenders() {
		// skip not equal
		if n != appenderName {
			continue
		}
		a.SetLevel(level, catalogs...)
	}
}
