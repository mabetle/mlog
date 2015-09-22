package wlog

import (
	"fmt"
	"strings"
)

// Appender define interface
type Appender interface {
	// BaseAppender vars
	GetName() string
	GetLevelMap() map[string]Level

	// BaseAppender has implements these.
	ScanConfigLevel(lines []string)
	IsOutputLog(level, catalog string) bool

	// each appender should implements WriteLog()
	WriteLog(level string, catalog string, callin int, msg ...interface{})
}

var Appenders = make(map[string]Appender)

// InitAppender
func GetAppenders() map[string]Appender {
	if len(Appenders) < 1 {
		AddAppender(NewConsoleAppender(), []string{})
	}
	return Appenders
}

// AddAppender
func AddAppender(appender Appender, lines []string) {
	appender.ScanConfigLevel(lines)
	Appenders[appender.GetName()] = appender
}

type BaseAppender struct {
	Name     string
	LevelMap map[string]Level
}

func NewBaseAppender(name string) *BaseAppender {
	m := &BaseAppender{}
	m.Name = name
	m.LevelMap = make(map[string]Level)
	return m
}

func (a *BaseAppender) GetName() string {
	return a.Name
}

func (a *BaseAppender) GetLevelMap() map[string]Level {
	return a.LevelMap
}

// Appender SetLevel
func (a *BaseAppender) SetLevel(levelLabel string, catalogs ...string) {
	level := GetStringLevel(levelLabel)
	if len(catalogs) == 0 {
		a.LevelMap[""] = level
		return
	}
	for _, catalog := range catalogs {
		//fmt.Printf("*****%s: %s:%v\n", a.Name, catalog, levelLabel)
		if catalog == "root" {
			catalog = ""
		}
		a.LevelMap[catalog] = level
	}
}

// ScanConfigLevel
func (a *BaseAppender) ScanConfigLevel(lines []string) {
	// process catelog level config
	levelPrefix := fmt.Sprintf("%s-level-", a.Name)
	//fmt.Printf("***%s\n", levelPrefix)
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

		a.SetLevel(level, catalog)
	}
}

// IsOutputLog
func (a *BaseAppender) IsOutputLog(levelLabel, catalog string) bool {
	level := GetStringLevel(levelLabel)
	appenderCatalogLevel, aok := GetCatalogLevel(catalog, a.GetLevelMap())
	catalogLevel, ok := GetCatalogLevel(catalog, levelMap)
	
	fmt.Printf("***LogLevel:%s, Appender Level: %v, %v CommonLevel: %v, %v \n", levelLabel, appenderCatalogLevel, aok, catalogLevel, ok)
	
	// 1.check appender config
	if aok && appenderCatalogLevel >= level {
		return true
	}
	// 2.check common level config
	if ok && catalogLevel >= level {
		return true
	}
	// 3.not found in config map
	if LevelInfo >= level {
		return true
	}
	// 4. less than INFO
	return false
}
