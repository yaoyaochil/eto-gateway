// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMemberInfoMileage = "member_info_mileage"

// MemberInfoMileage mapped from table <member_info_mileage>
type MemberInfoMileage struct {
	MID          int32     `gorm:"column:m_id;primaryKey;autoIncrement:true" json:"m_id"`
	UserID       string    `gorm:"column:user_id" json:"user_id"`
	UserName     string    `gorm:"column:user_name;not null" json:"user_name"`
	FirstSsn     string    `gorm:"column:first_ssn;not null" json:"first_ssn"`
	SecondSsn    string    `gorm:"column:second_ssn;not null" json:"second_ssn"`
	Passwd       string    `gorm:"column:passwd;not null" json:"passwd"`
	MobileNo     string    `gorm:"column:mobile_no;not null" json:"mobile_no"`
	RegDate      int32     `gorm:"column:reg_date;not null" json:"reg_date"`
	Email        string    `gorm:"column:email;not null" json:"email"`
	QNo          int32     `gorm:"column:q_no;not null" json:"q_no"`
	QAnswer      string    `gorm:"column:q_answer;not null" json:"q_answer"`
	UpdtDate     time.Time `gorm:"column:updt_date;not null;default:CURRENT_TIMESTAMP" json:"updt_date"`
	State        int32     `gorm:"column:state;not null;default:1" json:"state"`
	Nickname     string    `gorm:"column:nickname;not null" json:"nickname"`
	EmailYn      string    `gorm:"column:email_yn;not null;default:y" json:"email_yn"`
	SsnCheck     int32     `gorm:"column:ssn_check;not null" json:"ssn_check"`
	Slot         int32     `gorm:"column:slot;not null;default:8" json:"slot"`
	LastPlayTime time.Time `gorm:"column:last_play_time;not null" json:"last_play_time"`
	HangameFlag  int32     `gorm:"column:hangame_flag;not null" json:"hangame_flag"`
	HanmonFlag   int32     `gorm:"column:hanmon_flag;not null" json:"hanmon_flag"`
	Mileage      int32     `gorm:"column:mileage" json:"mileage"`
}

// TableName MemberInfoMileage's table name
func (*MemberInfoMileage) TableName() string {
	return TableNameMemberInfoMileage
}
