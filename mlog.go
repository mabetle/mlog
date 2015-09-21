package mlog

import (
	"mabetle/mlog/wlog"
)

// Set logger catalogs level.
func SetLevel(level string, catalog ...string) { wlog.SetLevel(level, catalog...) }

// shortcuts for set level
func SetTraceLevel(catalogs ...string) { SetLevel("trace", catalogs...) }
func SetDebugLevel(catalogs ...string) { SetLevel("debug", catalogs...) }
func SetInfoLevel(catalogs ...string)  { SetLevel("info", catalogs...) }
func SetWarnLevel(catalogs ...string)  { SetLevel("warn", catalogs...) }
func SetErrorLevel(catalogs ...string) { SetLevel("error", catalogs...) }
