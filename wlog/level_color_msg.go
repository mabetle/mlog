package wlog

import (
	"fmt"
	"github.com/agtorre/gocolorize"
	"github.com/mabetle/mcore/mterm"
)

var (
	trace = gocolorize.NewColor("magenta")
	debug = gocolorize.NewColor("cyan")
	info  = gocolorize.NewColor("green")
	warn  = gocolorize.NewColor("yellow")
	err   = gocolorize.NewColor("red")

	//helper functions to longen code
	d = debug.Paint
	i = info.Paint
	w = warn.Paint
	e = err.Paint
	t = trace.Paint
)

func GetLevelColorMsg(level string, v ...interface{}) string {
	// default no color.
	msg := fmt.Sprint(v...)
	levelLevel := GetStringLevel(level)

	// xterm can show color.
	if mterm.IsXterm() {
		switch levelLevel {
		case LevelInfo:
			msg = i(msg)
		case LevelTrace:
			msg = t(msg)
		case LevelError:
			msg = e(msg)
		case LevelWarn:
			msg = w(msg)
		case LevelDebug:
			msg = d(msg)
		default:
			msg = i(msg)
		}
	}
	return msg
}
