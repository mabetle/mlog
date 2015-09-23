package logapi

// Appender define interface
type Appender interface {
	// BaseAppender vars
	GetName() string

	// BaseAppender has implements these.
	SetLevel(level string, catalogs ...string)
	ScanConfigLevel(lines []string)
	IsOutputLog(level, catalog string) bool
	Inspect(catalog string)

	// each appender should implements WriteLog()
	WriteLog(level string, catalog string, callin int, msg ...interface{})
}
