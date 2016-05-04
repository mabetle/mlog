package logapi

// Logger interface
//
type Logger interface {
	// basic logger methods
	Trace(msg ...interface{})
	Debug(msg ...interface{})
	Info(msg ...interface{})
	Warn(msg ...interface{})
	Error(msg ...interface{})

	// format msg.
	Tracef(format string, msg ...interface{})
	Debugf(format string, msg ...interface{})
	Infof(format string, msg ...interface{})
	Warnf(format string, msg ...interface{})
	Errorf(format string, msg ...interface{})

	// check error with log
	CheckError(err error, msg ...interface{}) bool
	CheckErrorf(err error, msg string, args ...interface{}) bool

	//  set catalog level
	//SetLevel(level string, catalog ...string)

	// ouput log infomations
	Inspect() // show log info
}

type LogWriter interface {
	WriteLog(level string, catalog string, callin int, msg ...interface{})
}
