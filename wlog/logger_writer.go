package wlog

import (
	"strings"
)

// GetCatalogLevel
// if catalog level not matched, default to INFO
// you can change it by mlog.SetLevel()
func GetCatalogLevel(catalog string, levelMap map[string]Level) (l Level, ok bool) {
	// default level is info.
	for k, v := range levelMap {
		if len(k) > len(catalog) {
			continue
		}
		if strings.HasPrefix(catalog, k) {
			l = v
			ok = true
		}
	}
	//fmt.Printf("Catalog:%s Level:%s \n", catalog, GetLevelLabel(l))
	return
}

// WriteLog
// all log methods call this one.
func WriteLog(level, catalog string, callin int, args ...interface{}) {
	appenders := GetAppenders()
	for _, appender := range appenders {
		if !appender.IsOutputLog(catalog, level) {
			continue
		}
		appender.WriteLog(level, catalog, callin, args...)
	}
}
