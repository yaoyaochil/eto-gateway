// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAutoPunishHackIP = "auto_punish_hack_ip"

// AutoPunishHackIP mapped from table <auto_punish_hack_ip>
type AutoPunishHackIP struct {
	OccDate     time.Time `gorm:"column:occ_date;primaryKey" json:"occ_date"`
	HackType    int32     `gorm:"column:hack_type;primaryKey" json:"hack_type"`
	HackSubType int32     `gorm:"column:hack_sub_type;primaryKey" json:"hack_sub_type"`
	CClassIP    string    `gorm:"column:c_class_ip;primaryKey" json:"c_class_ip"`
	Cnt         int32     `gorm:"column:cnt;not null" json:"cnt"`
}

// TableName AutoPunishHackIP's table name
func (*AutoPunishHackIP) TableName() string {
	return TableNameAutoPunishHackIP
}
