# Mabetle Go Logger

Ideas from log4j.

## Feathers

* Support log level.
* Support catalog.
* Support appenders: Console, File, DB.
* Support color console output.
* Extendable, You can add new logger appender.
* Configable, using log.conf.

## Install
	
	go get -v github.com/mabetle/mlog

## Dependencies

Only depends on Go, no third part lib needed.

## Usage

### In Go code:

	import "github.com/mabetle/mlog"
	
	// xxx means logger catalog
	var logger = mlog.GetLogger("xxx")
	
	func init(){
		// set logger level
		mlog.SetLevel("info", "xxx")
		mlog.SetDebugLevel("xxx")
	}

	func XXX(){
		logger.Info("info")	
		logger.Infof("info %s", info)
	}

### Without configuration file, Init Logger in Go code:

	func init(){
		mlog.AddAppender(...)
		mlog.SetLevel(...)
	}

	var logger = mlog.GetLogger("xxx")

### Specific custom log config location:

	mlog.LoadConfig( location )

## Write Myself Appender

Appender interface define:

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

Your Appender implements can extends BaseAppender, BaseAppender provides all methods 
except WriteLog(...). You should write you own WriteLog method.


## Configuration

Mlog using log.conf to config logger runtime environment.

Search log.conf in following sequence:

* ./log.conf
* ./conf/log.conf
* /conf/log.conf
* /rundata/log.conf

Logger config file format reference: [misc/log_tml.conf](log_tml.conf)

## Known Bugs


## Improvement

Welcome to help improve this library.

Report BUGS:

* Start an issue.


Merge Request


## TODO

* Add rolling file support.


## License

[Apache License Version 2.0](LICENSE)



