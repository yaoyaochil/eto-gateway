// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTingUserSpec = "ting_user_spec"

// TingUserSpec mapped from table <ting_user_spec>
type TingUserSpec struct {
	MID         int32     `gorm:"column:m_id;primaryKey" json:"m_id"`
	RegDatetime time.Time `gorm:"column:reg_datetime;not null" json:"reg_datetime"`
	CPUVendor   string    `gorm:"column:cpu_vendor;not null;default:0" json:"cpu_vendor"`
	CPUNum      string    `gorm:"column:cpu_num;not null;default:0" json:"cpu_num"`
	CPUClock    int32     `gorm:"column:cpu_clock;not null" json:"cpu_clock"`
	RAM         int32     `gorm:"column:ram;not null" json:"ram"`
	VideoVendor int32     `gorm:"column:video_vendor;not null" json:"video_vendor"`
	VideoDevice int32     `gorm:"column:video_device;not null" json:"video_device"`
	VideoRAM    int32     `gorm:"column:video_ram;not null" json:"video_ram"`
	Os          string    `gorm:"column:os;not null;default:0" json:"os"`
	OsBit       string    `gorm:"column:os_bit;not null;default:0" json:"os_bit"`
}

// TableName TingUserSpec's table name
func (*TingUserSpec) TableName() string {
	return TableNameTingUserSpec
}