# Mabetle Go Logger

Ideas from log4j.

## Feathers

* Support log level.
* Support catalog.
* Support appenders: Console, File.
* Support color console output.
* Extendable, User can add new logger appender.

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

Config file:

reference: [log_tml.conf](log_tml.conf)

## Bugs


## TODO


## License

[Apache License Version 2.0](LICENSE)



