// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameUserBan = "user_ban"

// UserBan mapped from table <user_ban>
type UserBan struct {
	No           int32  `gorm:"column:no;primaryKey;autoIncrement:true" json:"no"`
	Category     int32  `gorm:"column:category;not null;default:1" json:"category"`
	MID          int32  `gorm:"column:m_id;not null" json:"m_id"`
	BanTerm      int32  `gorm:"column:ban_term;not null" json:"ban_term"`
	BanReason    int32  `gorm:"column:ban_reason;not null" json:"ban_reason"`
	DetailReason string `gorm:"column:detail_reason;not null" json:"detail_reason"`
	BanDate      int32  `gorm:"column:ban_date;not null" json:"ban_date"`
	CancelReason string `gorm:"column:cancel_reason;not null" json:"cancel_reason"`
	CancelDate   int32  `gorm:"column:cancel_date;not null" json:"cancel_date"`
	AdminID      int32  `gorm:"column:admin_id;not null" json:"admin_id"`
	Status       int32  `gorm:"column:status;not null" json:"status"`
	FirstSsn     string `gorm:"column:first_ssn;not null" json:"first_ssn"`
	SecondSsn    string `gorm:"column:second_ssn;not null" json:"second_ssn"`
}

// TableName UserBan's table name
func (*UserBan) TableName() string {
	return TableNameUserBan
}
