// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameEventGoldcardInfo = "event_goldcard_info"

// EventGoldcardInfo mapped from table <event_goldcard_info>
type EventGoldcardInfo struct {
	MID    int32 `gorm:"column:m_id;primaryKey" json:"m_id"`
	Coupon int32 `gorm:"column:coupon;not null" json:"coupon"`
}

// TableName EventGoldcardInfo's table name
func (*EventGoldcardInfo) TableName() string {
	return TableNameEventGoldcardInfo
}
