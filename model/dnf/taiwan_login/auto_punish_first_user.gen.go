// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAutoPunishFirstUser = "auto_punish_first_user"

// AutoPunishFirstUser mapped from table <auto_punish_first_user>
type AutoPunishFirstUser struct {
	MID         int32     `gorm:"column:m_id;primaryKey" json:"m_id"`
	OccTime     time.Time `gorm:"column:occ_time;not null" json:"occ_time"`
	IP          string    `gorm:"column:ip;not null" json:"ip"`
	HackType    int32     `gorm:"column:hack_type;primaryKey" json:"hack_type"`
	Cnt         int32     `gorm:"column:cnt;not null" json:"cnt"`
	PunishFlag  int32     `gorm:"column:punish_flag;not null" json:"punish_flag"`
	HackSubType int32     `gorm:"column:hack_sub_type;primaryKey" json:"hack_sub_type"`
	HackSubCnt  int32     `gorm:"column:hack_sub_cnt;not null" json:"hack_sub_cnt"`
}

// TableName AutoPunishFirstUser's table name
func (*AutoPunishFirstUser) TableName() string {
	return TableNameAutoPunishFirstUser
}
