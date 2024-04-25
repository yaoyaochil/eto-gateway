// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMemberSafeEnsure = "member_safe_ensure"

// MemberSafeEnsure mapped from table <member_safe_ensure>
type MemberSafeEnsure struct {
	MID         int32     `gorm:"column:m_id;not null" json:"m_id"`
	OccTime     time.Time `gorm:"column:occ_time;not null" json:"occ_time"`
	MobileNo    string    `gorm:"column:mobile_no;not null" json:"mobile_no"`
	ServiceFlag int32     `gorm:"column:service_flag;not null" json:"service_flag"`
	Type1Flag   int32     `gorm:"column:type1_flag;not null" json:"type1_flag"`
	Type2Flag   int32     `gorm:"column:type2_flag;not null" json:"type2_flag"`
	ExpireTime  time.Time `gorm:"column:expire_time;not null" json:"expire_time"`
	SettleID    string    `gorm:"column:settle_id;not null" json:"settle_id"`
}

// TableName MemberSafeEnsure's table name
func (*MemberSafeEnsure) TableName() string {
	return TableNameMemberSafeEnsure
}
