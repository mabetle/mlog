package wlog

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
