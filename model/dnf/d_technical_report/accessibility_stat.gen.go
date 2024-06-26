// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAccessibilityStat = "accessibility_stat"

// AccessibilityStat mapped from table <accessibility_stat>
type AccessibilityStat struct {
	OccTime  time.Time `gorm:"column:occ_time;primaryKey" json:"occ_time"`
	MainType int32     `gorm:"column:main_type;primaryKey" json:"main_type"`
	SubType  int32     `gorm:"column:sub_type;primaryKey" json:"sub_type"`
	Val      int32     `gorm:"column:val;not null" json:"val"`
}

// TableName AccessibilityStat's table name
func (*AccessibilityStat) TableName() string {
	return TableNameAccessibilityStat
}
