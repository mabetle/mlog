package wlog

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Default config location.
var ConfigLocation = "log.conf"

// LoadAppenders
func LoadAppenders(lines []string) {
	ScanAddConsoleAppender(lines)
	ScanAddFileAppender(lines)
	ScanAddDBAppender(lines)
}

// IsHasAppender
func IsHasAppender(name string, lines []string) bool {
	name = strings.ToLower(name)
	value, ok := ScanConfigValue("appender", lines)
	if !ok {
		return false
	}
	aS := strings.Split(value, ",")
	for _, a := range aS {
		a = strings.TrimSpace(a)
		a = strings.ToLower(a)
		if a == name {
			return true
		}
	}
	return false
}

// ScanConfigValue
func ScanConfigValue(scanKey string, lines []string) (string, bool) {
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
		// process append-file-location key
		if key == scanKey {
			return value, true
		}
	}
	return "", false
}

// LoadConfig try to read config file.
func LoadConfig(location string) error {
	bs, err := ioutil.ReadFile(location)
	if err != nil {
		// file not found
		return err
	}
	fmt.Printf("Init logger form: %s\n", location)
	lines := strings.Split(string(bs), "\n")

	LoadAppenders(lines)

	return nil
}

// LoadDefaultConfig loads logger configuations.
// Privilege sequence:
// - ConfigLocation specification
// - log.conf
// - conf/log.conf
// - /etc/mlog/log.conf
func InitConfig() {
	// Load config from ConfigLoaction
	if LoadConfig(ConfigLocation) == nil {
		return
	}

	// Load config from current dir
	if LoadConfig("log.conf") == nil {
		return
	}

	// Load config from conf dir
	if LoadConfig("conf/log.conf") == nil {
		return
	}

	// Load config from /etc
	if LoadConfig("/etc/mlog/log.conf") == nil {
		return
	}
}
