// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePowerwarLoading = "powerwar_loading"

// PowerwarLoading mapped from table <powerwar_loading>
type PowerwarLoading struct {
	MID          int32     `gorm:"column:m_id;primaryKey" json:"m_id"`
	OccTime      time.Time `gorm:"column:occ_time;primaryKey" json:"occ_time"`
	Round        int32     `gorm:"column:round;primaryKey" json:"round"`
	Player       int32     `gorm:"column:player;not null" json:"player"`
	MyLoading    int32     `gorm:"column:my_loading;not null" json:"my_loading"`
	OtherLoading int32     `gorm:"column:other_loading;not null" json:"other_loading"`
	VsLoading    int32     `gorm:"column:vs_loading;not null" json:"vs_loading"`
}

// TableName PowerwarLoading's table name
func (*PowerwarLoading) TableName() string {
	return TableNamePowerwarLoading
}