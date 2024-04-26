// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLogGiftHistory = "log_gift_history"

// LogGiftHistory gift history
type LogGiftHistory struct {
	TranID        int64     `gorm:"column:tran_id;primaryKey" json:"tran_id"`
	TranState     int32     `gorm:"column:tran_state;not null" json:"tran_state"`
	SendAccountID string    `gorm:"column:send_account_id;not null" json:"send_account_id"`
	SendCharacID  string    `gorm:"column:send_charac_id;not null" json:"send_charac_id"`
	RecvAccountID string    `gorm:"column:recv_account_id;not null" json:"recv_account_id"`
	ItemID        int32     `gorm:"column:item_id;not null" json:"item_id"`
	Cera          int32     `gorm:"column:cera;not null" json:"cera"`
	SendBeforCera int32     `gorm:"column:send_befor_cera;not null" json:"send_befor_cera"`
	SendAfterCera int32     `gorm:"column:send_after_cera;not null" json:"send_after_cera"`
	RecvBeforCera int32     `gorm:"column:recv_befor_cera;not null" json:"recv_befor_cera"`
	RecvAfterCera int32     `gorm:"column:recv_after_cera;not null" json:"recv_after_cera"`
	QueryUser     string    `gorm:"column:query_user;not null;default:None" json:"query_user"`
	OccDate       time.Time `gorm:"column:occ_date;not null" json:"occ_date"`
}

// TableName LogGiftHistory's table name
func (*LogGiftHistory) TableName() string {
	return TableNameLogGiftHistory
}
