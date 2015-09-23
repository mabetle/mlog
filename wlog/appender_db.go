package wlog

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"runtime"
	"sync"
	"time"
)

type DBAppender struct {
	mu  sync.Mutex // ensures atomic writes; protects the following fields
	Sql *sql.DB
	*BaseAppender
}

// NewDBAppender create new db appender
func NewDBAppender(driver, spec, table string) (*DBAppender, error) {
	if table == "" {
		table = "common_wlog"
	}
	DB_LOG_TABLE = table
	m := &DBAppender{}
	db, err := sql.Open(driver, spec)
	if err != nil {
		fmt.Printf("Error Open DB: %v", err)
		return nil, err
	}

	if err := Sync(db, LogTable{}); err != nil {
		// not show exists message.
		//fmt.Printf("Sync LogTable error: %s\n", err)
	}

	m.Sql = db
	m.BaseAppender = NewBaseAppender("db")
	return m, nil
}

func Sync(db *sql.DB, model LogTable) error {
	q := `
CREATE TABLE %s (
  Id bigint(20) NOT NULL AUTO_INCREMENT,
  Level varchar(10) DEFAULT NULL,
  Source varchar(255) DEFAULT NULL,
  Line int(11) DEFAULT NULL,
  Catalog varchar(500) DEFAULT NULL,
  Message varchar(1000) DEFAULT NULL,
  CreateTime datetime DEFAULT NULL,
  PRIMARY KEY (Id)
)`
	//ENGINE=InnoDB AUTO_INCREMENT=230 DEFAULT CHARSET=utf8
	q = fmt.Sprintf(q, DB_LOG_TABLE)
	if _, err := db.Exec(q); err != nil {
		return err
	}
	return nil
}

func Insert(db *sql.DB, model LogTable) error {
	q := `
INSERT INTO %s (Level,Source,Line,Catalog,Message,CreateTime) 
VALUES (?,?,?,?,?,?)
`
	q = fmt.Sprintf(q, DB_LOG_TABLE)
	if _, err := db.Exec(q, model.Level, model.Source, model.Line,
		model.Catalog, model.Message, model.CreateTime); err != nil {
		return err
	}
	return nil
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
	if err := Insert(a.Sql, m); err != nil {
		fmt.Printf("Write Logger to DB Error: %v\n", err)
	}

	a.mu.Unlock()
}

// AddDBAppender
func AddDBAppender(driver, spec string, table string, lines []string) {
	a, err := NewDBAppender(driver, spec, table)
	if err != nil {
		fmt.Printf("Error db appender: %v\n", err)
		return
	}
	AddAppender(a, lines)
}

// ScanAddDBAppender
func ScanAddDBAppender(lines []string) {
	if !IsHasAppender("db", lines) {
		return
	}
	driver, _ := ScanConfigValue("appender-db-driver", lines)
	spec, _ := ScanConfigValue("appender-db-spec", lines)
	table, _ := ScanConfigValue("appender-db-log-table", lines)
	// add db appender
	if driver != "" && spec != "" {
		AddDBAppender(driver, spec, table, lines)
	}
}
