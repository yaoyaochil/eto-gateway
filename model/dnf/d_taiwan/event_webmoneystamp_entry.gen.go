// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEventWebmoneystampEntry = "event_webmoneystamp_entry"

// EventWebmoneystampEntry mapped from table <event_webmoneystamp_entry>
type EventWebmoneystampEntry struct {
	MID            int32     `gorm:"column:m_id;primaryKey" json:"m_id"`
	OccTime        time.Time `gorm:"column:occ_time;not null" json:"occ_time"`
	AttendPoint    int32     `gorm:"column:attend_point;not null" json:"attend_point"`
	LastAttendTime time.Time `gorm:"column:last_attend_time;not null" json:"last_attend_time"`
	ReturnFlag     int32     `gorm:"column:return_flag;not null" json:"return_flag"`
	EntryItem      int32     `gorm:"column:entry_item;not null" json:"entry_item"`
}

// TableName EventWebmoneystampEntry's table name
func (*EventWebmoneystampEntry) TableName() string {
	return TableNameEventWebmoneystampEntry
}
