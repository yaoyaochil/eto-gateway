// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePowerWar = "power_war"

// PowerWar mapped from table <power_war>
type PowerWar struct {
	ServerID   int32     `gorm:"column:server_id;primaryKey" json:"server_id"`
	ASidePoint int32     `gorm:"column:a_side_point;not null" json:"a_side_point"`
	BSidePoint int32     `gorm:"column:b_side_point;not null" json:"b_side_point"`
	WinnerSide int32     `gorm:"column:winner_side;not null" json:"winner_side"`
	OccTime    time.Time `gorm:"column:occ_time;not null" json:"occ_time"`
}

// TableName PowerWar's table name
func (*PowerWar) TableName() string {
	return TableNamePowerWar
}
