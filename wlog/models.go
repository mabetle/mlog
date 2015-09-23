package wlog

import (
	"time"
)

// You can change it outside
var DB_LOG_TABLE = "common_wlog"

// store in database
type LogTable struct {
	Id         int64  ``
	Level      string `xorm:"varchar(10)"`
	Source     string `xorm:"varchr(500)"`
	Line       int
	Catalog    string    `xorm:"varchar(500)"`
	Message    string    `xorm:"varchar(1000)"`
	CreateTime time.Time `xorm:"created"`
}

// TableName
func (m LogTable) TableName() string {
	return DB_LOG_TABLE
}
