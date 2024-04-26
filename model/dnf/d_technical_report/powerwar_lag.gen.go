// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePowerwarLag = "powerwar_lag"

// PowerwarLag mapped from table <powerwar_lag>
type PowerwarLag struct {
	MID     int32     `gorm:"column:m_id;primaryKey" json:"m_id"`
	OccTime time.Time `gorm:"column:occ_time;primaryKey" json:"occ_time"`
	Round   int32     `gorm:"column:round;primaryKey" json:"round"`
	Player  int32     `gorm:"column:player;not null" json:"player"`
	LagAvg  float32   `gorm:"column:lag_avg;not null" json:"lag_avg"`
	LagCnt  float32   `gorm:"column:lag_cnt;not null" json:"lag_cnt"`
}

// TableName PowerwarLag's table name
func (*PowerwarLag) TableName() string {
	return TableNamePowerwarLag
}
