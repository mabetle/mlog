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
	//mlog.SetInfoLevel()
	mlog.SetErrorLevel()
	//mlogApiLogger.Error("Hello")
	mlogWrapLogger.Trace("Hello")
	mlogWrapLogger.Debug("Hello")
	mlogWrapLogger.Info("Hello")
	mlogWrapLogger.Error("Hello")
	mlogWrapLogger.Warn("Hello")

	//mlogWrapLogger.Inspect()

	//wlogApiLogger.Error("Hello")
	//wlogWrapLogger.Error("Hello")
}

func main() {
	Demo()
}
