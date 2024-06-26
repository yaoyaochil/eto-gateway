// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAutoPunishSecondLog = "auto_punish_second_log"

// AutoPunishSecondLog mapped from table <auto_punish_second_log>
type AutoPunishSecondLog struct {
	MID       int32     `gorm:"column:m_id;primaryKey" json:"m_id"`
	HackMID   int32     `gorm:"column:hack_m_id;primaryKey" json:"hack_m_id"`
	OccTime   time.Time `gorm:"column:occ_time;not null" json:"occ_time"`
	TradeCnt  int32     `gorm:"column:trade_cnt;not null" json:"trade_cnt"`
	TradeGold int64     `gorm:"column:trade_gold;not null" json:"trade_gold"`
}

// TableName AutoPunishSecondLog's table name
func (*AutoPunishSecondLog) TableName() string {
	return TableNameAutoPunishSecondLog
}
