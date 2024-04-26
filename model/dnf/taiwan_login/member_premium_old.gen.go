// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMemberPremiumOld = "member_premium_old"

// MemberPremiumOld mapped from table <member_premium_old>
type MemberPremiumOld struct {
	EventID      int32     `gorm:"column:event_id;primaryKey" json:"event_id"`
	PreType      int32     `gorm:"column:pre_type;primaryKey" json:"pre_type"`
	MID          int32     `gorm:"column:m_id;primaryKey" json:"m_id"`
	ServiceStart time.Time `gorm:"column:service_start;primaryKey" json:"service_start"`
	ServiceEnd   time.Time `gorm:"column:service_end;not null" json:"service_end"`
	ServerID     int32     `gorm:"column:server_id;primaryKey" json:"server_id"`
}

// TableName MemberPremiumOld's table name
func (*MemberPremiumOld) TableName() string {
	return TableNameMemberPremiumOld
}
