// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLoginAccount1 = "login_account_1"

// LoginAccount1 mapped from table <login_account_1>
type LoginAccount1 struct {
	MID           int32     `gorm:"column:m_id;primaryKey" json:"m_id"`
	MChannelNo    int32     `gorm:"column:m_channel_no;not null" json:"m_channel_no"`
	LoginStatus   bool      `gorm:"column:login_status;not null" json:"login_status"`
	LastLoginDate time.Time `gorm:"column:last_login_date;not null" json:"last_login_date"`
	LoginIP       string    `gorm:"column:login_ip;not null" json:"login_ip"`
}

// TableName LoginAccount1's table name
func (*LoginAccount1) TableName() string {
	return TableNameLoginAccount1
}
