package wlog

import (
	"log"
)

var (
	//Define the look/feel of the INFO logger
	//LogFormat = log.Ldate | log.Ltime | log.Llongfile
	LogFormat = log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile
	//LogFormat = log.Llongfile
	//LogFormat = log.Lshortfile
)
