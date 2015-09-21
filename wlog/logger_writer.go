package wlog

import (
	"strings"
)

// GetCatalogLevel
// if catalog level not matched, default to INFO
// you can change it by mlog.SetLevel()
func GetCatalogLevel(catalog string) (l Level) {
	// default level is info.
	l = LevelInfo
	for k, v := range levelMap {
		if len(k) > len(catalog) {
			continue
		}
		if strings.HasPrefix(catalog, k) {
			l = v
		}
	}
	//fmt.Printf("Catalog:%s Level:%s \n", catalog, GetLevelLabel(l))
	return
}

// when level big than catalog level set, then output.
func IsOutputLog(catalog string, level string) bool {
	catalogLevel := GetCatalogLevel(catalog)
	levelLevel := GetStringLevel(level)
	return levelLevel >= catalogLevel
}

// WriteLog
// all log methods call this one.
func WriteLog(level, catalog string, callin int, args ...interface{}) {
	if !IsOutputLog(catalog, level) {
		return
	}
	// set log catalog
	if catalog == "" {
		catalog = "default"
	}

	appenders := GetAppenders()

	for _, appender := range appenders {
		appender.WriteLog(level, catalog, callin, args...)
	}
}
