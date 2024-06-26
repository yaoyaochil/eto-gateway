// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLogRechargeHistory = "log_recharge_history"

// LogRechargeHistory recharge history
type LogRechargeHistory struct {
	TranID      int64     `gorm:"column:tran_id;primaryKey" json:"tran_id"`
	OrderTranID string    `gorm:"column:order_tran_id;not null" json:"order_tran_id"`
	TranState   int32     `gorm:"column:tran_state;not null" json:"tran_state"`
	AccountID   string    `gorm:"column:account_id;not null" json:"account_id"`
	CharacID    string    `gorm:"column:charac_id;not null" json:"charac_id"`
	Cera        int32     `gorm:"column:cera;not null" json:"cera"`
	BeforCera   int32     `gorm:"column:befor_cera;not null" json:"befor_cera"`
	AfterCera   int32     `gorm:"column:after_cera;not null" json:"after_cera"`
	ChargeType  int32     `gorm:"column:charge_type;not null" json:"charge_type"`
	QueryUser   string    `gorm:"column:query_user;not null;default:None" json:"query_user"`
	OccDate     time.Time `gorm:"column:occ_date;not null" json:"occ_date"`
}

// TableName LogRechargeHistory's table name
func (*LogRechargeHistory) TableName() string {
	return TableNameLogRechargeHistory
}
