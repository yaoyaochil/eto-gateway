// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMemberMacInfo = "member_mac_info"

// MemberMacInfo mapped from table <member_mac_info>
type MemberMacInfo struct {
	No      int32     `gorm:"column:no;primaryKey;autoIncrement:true" json:"no"`
	MacAddr string    `gorm:"column:mac_addr;not null" json:"mac_addr"`
	OccTime time.Time `gorm:"column:occ_time;not null" json:"occ_time"`
}

// TableName MemberMacInfo's table name
func (*MemberMacInfo) TableName() string {
	return TableNameMemberMacInfo
}
