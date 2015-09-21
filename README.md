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


## Usage

In Go code:

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

## Configuration

Mlog using log.conf to config logger runtime environment.

Search log.conf in following sequence:

* ./log.conf
* ./conf/log.conf
* /conf/log.conf
* /rundata/log.conf

Logger config file format reference: [log_tml.conf](log_tml.conf)

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



