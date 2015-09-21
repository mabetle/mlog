package wlog

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"sync"
	"runtime"
)

// store in database
type LogTable struct {
	Id         int64    ``
	Level      string `xorm:"varchar(10)"`
	Source     string `xorm:"varchr(500)"`
	Line       int
	Catalog    string `xorm:"varchar(500)"`
	Message    string `xorm:"varchar(1000)"`
	CreateTime time.Time
}

// TableName
func (m LogTable) TableName() string {
	return "common_wlog"
}

type DBAppender struct {
	mu     sync.Mutex // ensures atomic writes; protects the following fields
	Xorm *xorm.Engine
}

// NewDBAppender create new db appender
func NewDBAppender(driver, spec string) (*DBAppender, error) {
	m := &DBAppender{}
	e, err := xorm.NewEngine(driver, spec)
	if err != nil {
		return nil, err
	}
	// Sync Table
	e.Sync(LogTable{})

	m.Xorm = e
	return m, nil
}

// WriteLog implements for api.LogWriter
func (a *DBAppender) WriteLog(level string, catalog string, callin int, v ...interface{}) {

	// arg call in not used.
	callin = 3
	a.mu.Lock()
	m := LogTable{}
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

func AddDBAppender(driver, spec string) {
	a, err := NewDBAppender(driver, spec)
	if err != nil {
		fmt.Printf("Error db appender: %v\n", err)
		return
	}

	Appenders["db"] = a
}
