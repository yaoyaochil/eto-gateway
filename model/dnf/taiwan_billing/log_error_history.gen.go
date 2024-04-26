// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLogErrorHistory = "log_error_history"

// LogErrorHistory mapped from table <log_error_history>
type LogErrorHistory struct {
	No         int32     `gorm:"column:no;primaryKey;autoIncrement:true" json:"no"`
	ErrorID    int32     `gorm:"column:error_id;not null" json:"error_id"`
	ErrorMsg   string    `gorm:"column:error_msg;not null" json:"error_msg"`
	ErrorQuery string    `gorm:"column:error_query;not null" json:"error_query"`
	ProcName   string    `gorm:"column:proc_name;not null" json:"proc_name"`
	ProcLine   int32     `gorm:"column:proc_line;not null" json:"proc_line"`
	QueryUser  string    `gorm:"column:query_user;not null;default:None" json:"query_user"`
	OccDate    time.Time `gorm:"column:occ_date;not null" json:"occ_date"`
}

// TableName LogErrorHistory's table name
func (*LogErrorHistory) TableName() string {
	return TableNameLogErrorHistory
}
