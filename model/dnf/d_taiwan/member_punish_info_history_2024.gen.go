// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMemberPunishInfoHistory2024 = "member_punish_info_history_2024"

// MemberPunishInfoHistory2024 mapped from table <member_punish_info_history_2024>
type MemberPunishInfoHistory2024 struct {
	No          int32     `gorm:"column:no;primaryKey;autoIncrement:true" json:"no"`
	MID         int32     `gorm:"column:m_id;not null" json:"m_id"`
	PunishType  int32     `gorm:"column:punish_type;not null" json:"punish_type"`
	OccTime     time.Time `gorm:"column:occ_time;not null" json:"occ_time"`
	PunishValue int32     `gorm:"column:punish_value;not null" json:"punish_value"`
	ApplyFlag   int32     `gorm:"column:apply_flag;not null" json:"apply_flag"`
	StartTime   time.Time `gorm:"column:start_time;not null" json:"start_time"`
	EndTime     time.Time `gorm:"column:end_time;not null" json:"end_time"`
	AdminID     string    `gorm:"column:admin_id" json:"admin_id"`
	Reason      string    `gorm:"column:reason" json:"reason"`
	IsKicked    int32     `gorm:"column:is_kicked" json:"is_kicked"`
	FirstSsn    string    `gorm:"column:first_ssn" json:"first_ssn"`
	SecondSsn   string    `gorm:"column:second_ssn" json:"second_ssn"`
}

// TableName MemberPunishInfoHistory2024's table name
func (*MemberPunishInfoHistory2024) TableName() string {
	return TableNameMemberPunishInfoHistory2024
}
