// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserTingTimecheck = "user_ting_timecheck"

// UserTingTimecheck mapped from table <user_ting_timecheck>
type UserTingTimecheck struct {
	OccTime time.Time `gorm:"column:occ_time;primaryKey" json:"occ_time"`
	Minute  int32     `gorm:"column:minute;primaryKey" json:"minute"`
	Cnt     int32     `gorm:"column:cnt;not null" json:"cnt"`
}

// TableName UserTingTimecheck's table name
func (*UserTingTimecheck) TableName() string {
	return TableNameUserTingTimecheck
}
