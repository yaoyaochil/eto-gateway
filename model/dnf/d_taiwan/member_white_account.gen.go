// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMemberWhiteAccount = "member_white_account"

// MemberWhiteAccount mapped from table <member_white_account>
type MemberWhiteAccount struct {
	MID     int32     `gorm:"column:m_id;primaryKey" json:"m_id"`
	RegDate time.Time `gorm:"column:reg_date;not null" json:"reg_date"`
}

// TableName MemberWhiteAccount's table name
func (*MemberWhiteAccount) TableName() string {
	return TableNameMemberWhiteAccount
}
