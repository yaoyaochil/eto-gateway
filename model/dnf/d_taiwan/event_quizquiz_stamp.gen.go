// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameEventQuizquizStamp = "event_quizquiz_stamp"

// EventQuizquizStamp mapped from table <event_quizquiz_stamp>
type EventQuizquizStamp struct {
	MID     int32     `gorm:"column:m_id;primaryKey" json:"m_id"`
	Degree  int32     `gorm:"column:degree;primaryKey" json:"degree"`
	Stamp   int32     `gorm:"column:stamp;not null" json:"stamp"`
	OccTime time.Time `gorm:"column:occ_time;not null" json:"occ_time"`
}

// TableName EventQuizquizStamp's table name
func (*EventQuizquizStamp) TableName() string {
	return TableNameEventQuizquizStamp
}
