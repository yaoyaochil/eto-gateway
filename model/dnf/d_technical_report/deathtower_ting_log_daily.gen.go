// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameDeathtowerTingLogDaily = "deathtower_ting_log_daily"

// DeathtowerTingLogDaily mapped from table <deathtower_ting_log_daily>
type DeathtowerTingLogDaily struct {
	OccDate time.Time `gorm:"column:occ_date;primaryKey" json:"occ_date"`
	Level   int32     `gorm:"column:level;primaryKey" json:"level"`
	TingCnt int32     `gorm:"column:ting_cnt;not null" json:"ting_cnt"`
}

// TableName DeathtowerTingLogDaily's table name
func (*DeathtowerTingLogDaily) TableName() string {
	return TableNameDeathtowerTingLogDaily
}
