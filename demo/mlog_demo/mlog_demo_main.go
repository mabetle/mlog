package main

import (
	"github.com/mabetle/mlog"
	"github.com/mabetle/mlog/wlog"
)

var (
	mlogApiLogger  = mlog.GetLogger("mlogApiLogger")         // api
	mlogWrapLogger = mlog.GetWrapLogger("mlogWrapLogger")    // mlog
	wlogApiLogger  = wlog.NewLogger("wlogApiLogger", 5)      // api
	wlogWrapLogger = wlog.NewWrapLogger("wlogWrapLogger", 5) //
)

func Demo() {
	mlogWrapLogger.Trace("Hello")
	mlogWrapLogger.Debug("Hello")
	mlogWrapLogger.Info("Hello")
	mlogWrapLogger.Warn("Hello")
	mlogWrapLogger.Error("Hello")

	mlogWrapLogger.Inspect()

	//wlogApiLogger.Error("Hello")
	//wlogWrapLogger.Error("Hello")
}

func main() {
	//mlog.SetTraceLevel()
	mlog.SetDebugLevel()
	//mlog.SetInfoLevel()
	//mlog.SetWarnLevel()
	//mlog.SetErrorLevel()
	Demo()
}
