package wlog

import (
	"strings"
)

// GetCatalogLevel
// if catalog level not matched, default to INFO
// you can change it by mlog.SetLevel()
func GetCatalogLevel(catalog string, lm map[string]Level) (l Level, ok bool) {
	for k, v := range lm {
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
func WriteLog(levelLabel, catalog string, callin int, args ...interface{}) {
	appenders := GetAppenders()
	for _, appender := range appenders {
		if !appender.IsOutputLog(levelLabel, catalog) {
			continue
		}
		appender.WriteLog(levelLabel, catalog, callin, args...)
	}
}
