// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAutoPunishBlackipInfo = "auto_punish_blackip_info"

// AutoPunishBlackipInfo mapped from table <auto_punish_blackip_info>
type AutoPunishBlackipInfo struct {
	IP        string    `gorm:"column:ip;primaryKey" json:"ip"`
	StartIP   int32     `gorm:"column:start_ip;primaryKey" json:"start_ip"`
	EndIP     int32     `gorm:"column:end_ip;primaryKey" json:"end_ip"`
	RegDate   time.Time `gorm:"column:reg_date;not null" json:"reg_date"`
	ApplyFlag int32     `gorm:"column:apply_flag;not null" json:"apply_flag"`
}

// TableName AutoPunishBlackipInfo's table name
func (*AutoPunishBlackipInfo) TableName() string {
	return TableNameAutoPunishBlackipInfo
}
