// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLogRefundHistory = "log_refund_history"

// LogRefundHistory mapped from table <log_refund_history>
type LogRefundHistory struct {
	TranID      int64     `gorm:"column:tran_id;primaryKey" json:"tran_id"`
	AccountID   string    `gorm:"column:account_id;primaryKey" json:"account_id"`
	OrderTranID string    `gorm:"column:order_tran_id;not null" json:"order_tran_id"`
	Amount      int32     `gorm:"column:amount;not null" json:"amount"`
	TranState   int32     `gorm:"column:tran_state;not null" json:"tran_state"`
	QueryUser   string    `gorm:"column:query_user;not null;default:None" json:"query_user"`
	OccDate     time.Time `gorm:"column:occ_date;not null" json:"occ_date"`
}

// TableName LogRefundHistory's table name
func (*LogRefundHistory) TableName() string {
	return TableNameLogRefundHistory
}
