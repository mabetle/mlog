package wlog

import (
	"strings"
)

type Level int

const (
	LevelTrace Level = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)

var levelMap = make(map[string]Level)

var levelLabelMap = make(map[Level]string)
var labelLevelMap = make(map[string]Level)

func InitLevel() {
	// init level label map
	levelLabelMap[LevelDebug] = "DEBUG"
	levelLabelMap[LevelWarn] = " WARN"
	levelLabelMap[LevelError] = "ERROR"
	levelLabelMap[LevelTrace] = "TRACE"
	levelLabelMap[LevelInfo] = " INFO"

	// init label level map
	labelLevelMap["INFO"] = LevelInfo
	labelLevelMap["WARN"] = LevelWarn
	labelLevelMap["ERROR"] = LevelError
	labelLevelMap["TRACE"] = LevelTrace
	labelLevelMap["DEBUG"] = LevelDebug
}

func GetLevelLabel(level Level) string {
	if v, ok := levelLabelMap[level]; ok {
		return v
	}
	return "INFO"
}

func GetStringLevel(str string) Level {
	str = strings.ToUpper(str)
	str = strings.TrimSpace(str)
	// if not found, set to INFO
	if v, ok := labelLevelMap[str]; ok {
		return v
	}
	return LevelInfo
}

// SetLevel implements api
func SetLevel(level string, catalogs ...string) {
	SetCatalogsLevel(GetStringLevel(level), catalogs...)
}

// SetCatalogsLevel is the primary function to set logger level.
func SetCatalogsLevel(level Level, catalogs ...string) {
	// not input catalog, set root
	if len(catalogs) == 0 {
		levelMap[""] = level
		return
	}

	for _, catalog := range catalogs {
		if catalog == "root" {
			catalog = ""
		}
		levelMap[catalog] = level
	}
}
