// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLoadingTime = "loading_time"

// LoadingTime mapped from table <loading_time>
type LoadingTime struct {
	OccTime  time.Time `gorm:"column:occ_time;primaryKey" json:"occ_time"`
	ServerID int32     `gorm:"column:server_id;primaryKey" json:"server_id"`
	Type     int32     `gorm:"column:type;primaryKey" json:"type"`
	LoadSec  int32     `gorm:"column:load_sec;not null" json:"load_sec"`
}

// TableName LoadingTime's table name
func (*LoadingTime) TableName() string {
	return TableNameLoadingTime
}