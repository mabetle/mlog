package wlog

// Appender define interface
type Appender interface {
	WriteLog(level string, catalog string, callin int, msg ...interface{})
}

var Appenders = make(map[string]Appender)

// InitAppender
func GetAppenders() map[string]Appender {
	if len(Appenders) < 1 {
		Appenders["Console"] = NewConsoleAppender()
	}
	return Appenders
}

// AddAppender
func AddAppender(name string, appender Appender) {
	Appenders[name] = appender
}
