// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGuildEvent = "guild_event"

// GuildEvent mapped from table <guild_event>
type GuildEvent struct {
	Gno     int32     `gorm:"column:gno;primaryKey" json:"gno"`
	SttDate time.Time `gorm:"column:stt_date;not null" json:"stt_date"`
	EndDate time.Time `gorm:"column:end_date;not null" json:"end_date"`
	AnnDate time.Time `gorm:"column:ann_date;not null" json:"ann_date"`
	PageURL string    `gorm:"column:page_url;not null" json:"page_url"`
}

// TableName GuildEvent's table name
func (*GuildEvent) TableName() string {
	return TableNameGuildEvent
}