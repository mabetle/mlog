package xorm_impl

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/mabetle/mlog/wlog"
	"runtime"
	"sync"
	"time"
)

type XormDBAppender struct {
	mu   sync.Mutex // ensures atomic writes; protects the following fields
	Xorm *xorm.Engine
	*wlog.BaseAppender
}

// NewXormDBAppender create new db appender
func NewXormDBAppender(driver, spec, table string) (*XormDBAppender, error) {
	if table == "" {
		table = "common_wlog"
	}

	wlog.DB_LOG_TABLE = table

	m := &XormDBAppender{}

	e, err := xorm.NewEngine(driver, spec)
	if err != nil {
		return nil, err
	}

	// config Engine
	e.SetTableMapper(core.SnakeMapper{})
	e.SetColumnMapper(core.SameMapper{})

	// Sync Table
	e.Sync(wlog.LogTable{})

	m.Xorm = e

	m.BaseAppender = wlog.NewBaseAppender("db")

	return m, nil
}

// WriteLog implements for api.LogWriter
func (a *XormDBAppender) WriteLog(level string, catalog string, callin int, v ...interface{}) {

	// arg call in not used.
	callin = 3
	a.mu.Lock()
	m := wlog.LogTable{}
	m.Level = level
	m.Catalog = catalog
	m.Message = fmt.Sprint(v...)
	m.CreateTime = time.Now()

	_, file, line, ok := runtime.Caller(callin)
	if !ok {
		file = "???"
		line = 0
	}

	m.Source = file
	m.Line = line

	// do save
	_, err := a.Xorm.Insert(&m)

	a.mu.Unlock()

	if err != nil {
		fmt.Printf("Write Logger to DB Error: %v\n", err)
	}
}

// AddXormDBAppender
func AddXormDBAppender(driver, spec string, table string, lines []string) {

	a, err := NewXormDBAppender(driver, spec, table)

	if err != nil {
		fmt.Printf("Error db appender: %v\n", err)
		return
	}
	wlog.AddAppender(a, lines)
}

// ScanAddXormDBAppender
func ScanAddXormDBAppender(lines []string) {
	if !wlog.IsHasAppender("db", lines) {
		return
	}
	driver, _ := wlog.ScanConfigValue("appender-db-driver", lines)
	spec, _ := wlog.ScanConfigValue("appender-db-spec", lines)
	table, _ := wlog.ScanConfigValue("appender-db-log-table", lines)
	// add db appender
	if driver != "" && spec != "" {
		AddXormDBAppender(driver, spec, table, lines)
	}
}
