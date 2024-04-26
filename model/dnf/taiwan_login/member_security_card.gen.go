// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMemberSecurityCard = "member_security_card"

// MemberSecurityCard mapped from table <member_security_card>
type MemberSecurityCard struct {
	MID           int32     `gorm:"column:m_id;primaryKey" json:"m_id"`
	OccTime       time.Time `gorm:"column:occ_time;not null" json:"occ_time"`
	Phone         string    `gorm:"column:phone;not null" json:"phone"`
	CertKey       string    `gorm:"column:cert_key;not null" json:"cert_key"`
	ServerKey     string    `gorm:"column:server_key;not null" json:"server_key"`
	Card          string    `gorm:"column:card;not null" json:"card"`
	FailCnt       int32     `gorm:"column:fail_cnt;not null" json:"fail_cnt"`
	ReIssueCnt    int32     `gorm:"column:re_issue_cnt;not null" json:"re_issue_cnt"`
	LastIssueTime time.Time `gorm:"column:last_issue_time;not null" json:"last_issue_time"`
	ValidityTime  int32     `gorm:"column:validity_time;not null" json:"validity_time"`
	ApplyFlag     int32     `gorm:"column:apply_flag;not null" json:"apply_flag"`
	CancelCnt     int32     `gorm:"column:cancel_cnt;not null" json:"cancel_cnt"`
	WebFlag       int32     `gorm:"column:web_flag;not null" json:"web_flag"`
	CertFlag      string    `gorm:"column:cert_flag;not null;default:0" json:"cert_flag"`
}

// TableName MemberSecurityCard's table name
func (*MemberSecurityCard) TableName() string {
	return TableNameMemberSecurityCard
}