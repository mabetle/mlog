package wlog

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Default config location.
var ConfigLocation = "log.conf"

func LoadAppenders(lines []string) {
	var hasConsoleAppender, hasFileAppender bool
	var fileAppendLocation string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// skip blank and comment line
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		lineA := strings.Split(line, "=")
		// skip not include = line
		if len(lineA) != 2 {
			continue
		}
		key := strings.TrimSpace(lineA[0])
		key = strings.ToLower(key)
		value := strings.TrimSpace(lineA[1])
		// process appender key
		if key == "appender" {
			aS := strings.Split(value, ",")
			for _, a := range aS {
				a = strings.TrimSpace(a)
				a = strings.ToLower(a)
				switch a {
				case "console":
					hasConsoleAppender = true
				case "file":
					hasFileAppender = true
				default:
				}
			}

		}
		// process append-file-location key
		if key == "appender-file-location" {
			fileAppendLocation = expandPath(value)
		}
	}
	// init appender
	if hasConsoleAppender {
		AddConsoleAppender()
	}

	if hasFileAppender && fileAppendLocation != "" {
		//fmt.Printf("fileAppendLocation: %s\n",fileAppendLocation )
		AddFileAppender(fileAppendLocation)
	}
}

// LoadConfig try to read config file.
func LoadConfig(location string) error {
	bs, err := ioutil.ReadFile(location)
	if err != nil {
		// file not found
		return err
	}

	lines := strings.Split(string(bs), "\n")

	LoadAppenders(lines)
	LoadConfigLevels(lines)

	return nil
}

// LoadConfigLevels
func LoadConfigLevels(lines []string) {
	// process catelog level config
	levelPrefix := "level-"

	for _, line := range lines {
		line = strings.TrimSpace(line)
		// skip blank line and comment line
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// skip not start with level-
		if !strings.HasPrefix(line, levelPrefix) {
			continue
		}

		kv := strings.Split(line, "=")

		if len(kv) != 2 {
			continue
		}

		catalog := strings.TrimPrefix(kv[0], levelPrefix)
		level := kv[1]
		SetLevel(level, catalog)
	}
}

// LoadDefaultConfig loads logger configuations.
// Privilege sequence:
// - ConfigLocation specification
// - log.conf
// - conf/log.conf
// - /rundata/log.conf
// - /conf/common/log.conf
func InitConfig() {
	// Load config from ConfigLoaction
	if LoadConfig(ConfigLocation) == nil {
		fmt.Printf("Init logger form: %s\n", ConfigLocation)
		return
	}

	// Load config from current dir
	if LoadConfig("log.conf") == nil {
		fmt.Printf("Init logger form: %s\n", "log.conf")
		return
	}

	// Load config from conf dir
	if LoadConfig("conf/log.conf") == nil {
		fmt.Printf("Init logger form: %s\n", "conf/log.conf")
		return
	}

	// Load config from /rundata dir
	if LoadConfig("/rundata/log.conf") == nil {
		fmt.Printf("Init logger form: %s\n", "/rundata/log.conf")
		return
	}

	// Load config from /conf dir
	if LoadConfig("/conf/common/log.conf") == nil {
		fmt.Printf("Init logger form: %s\n", "/conf/log.conf")
		return
	}

}
