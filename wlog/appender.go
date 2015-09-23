package wlog

import (
	"fmt"
	"github.com/mabetle/mlog/logapi"
	"strings"
)

var appenderMap = make(map[string]logapi.Appender)

// InitAppender
func GetAppenders() map[string]logapi.Appender {
	if len(appenderMap) < 1 {
		AddAppender(NewConsoleAppender(), []string{})
	}
	return appenderMap
}

// AddAppender
func AddAppender(appender logapi.Appender, lines []string) {
	// scan config level first
	appender.ScanConfigLevel(lines)

	appenderMap[appender.GetName()] = appender
}

// BaseAppender
type BaseAppender struct {
	Name     string
	levelMap map[string]Level
}

func NewBaseAppender(name string) *BaseAppender {
	m := &BaseAppender{}
	m.Name = name
	m.levelMap = make(map[string]Level)
	return m
}

func (a *BaseAppender) GetName() string {
	return a.Name
}

// Appender SetLevel
func (a *BaseAppender) SetLevel(levelLabel string, catalogs ...string) {
	level := GetLabelLevel(levelLabel)
	// SetRootLevel
	if len(catalogs) == 0 {
		a.levelMap[""] = level
		return
	}

	for _, catalog := range catalogs {
		if catalog == "root" {
			catalog = ""
		}
		a.levelMap[catalog] = level
	}
}

// ScanConfigLevel
func (a *BaseAppender) ScanConfigLevel(lines []string) {
	// process catelog level config
	a.scanPrefixConfigLevel("level-", lines)

	levelPrefix := fmt.Sprintf("%s-level-", a.Name)
	a.scanPrefixConfigLevel(levelPrefix, lines)
}

func (a *BaseAppender) scanPrefixConfigLevel(levelPrefix string, lines []string) {

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

// GetCatalogLevel
func (a *BaseAppender) getCatalogLevel(catalog string) Level {
	// found
	if l, ok := a.levelMap[catalog]; ok {
		return l
	}

	matchLen := 0
	cLen := len(catalog)

	// default not found set LevelInfo
	l := LevelInfo
	// set level to root level if found.
	if dl, ok := a.levelMap[""]; ok {
		l = dl
	}

	for k, v := range a.levelMap {
		kLen := len(k)
		// greater
		if kLen > cLen {
			continue
		}
		// equal
		if k == catalog {
			return v
		}
		// less
		// check "xxxx/xxx" k: xxxx
		if strings.HasPrefix(catalog, k) {
			if matchLen < kLen {
				matchLen = kLen
				l = v
			}
		}
	}
	return l
}

func (a *BaseAppender) IsOutputLog(label, catalog string) bool {
	checkLevel := GetLabelLevel(label)
	catalogLevel := a.getCatalogLevel(catalog)
	// error > info, true
	if checkLevel >= catalogLevel {
		return true
	}
	return false
}

func (a *BaseAppender) Inspect(catalog string) {
	fmt.Printf("\nAppender: %s\n", a.Name)

	fmt.Printf("\tLevel Map:\n")
	for c, l := range a.levelMap {
		fmt.Printf("\t\t%s:%s\n", c, GetLevelLabel(l))
	}

	cl := a.getCatalogLevel(catalog)
	fmt.Printf("Catalog: %s Level: %s \n", catalog, GetLevelLabel(cl))
}
