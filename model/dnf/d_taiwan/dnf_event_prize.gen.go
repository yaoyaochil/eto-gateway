// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDnfEventPrize = "dnf_event_prize"

// DnfEventPrize mapped from table <dnf_event_prize>
type DnfEventPrize struct {
	PrizeID   int32 `gorm:"column:prize_id;primaryKey" json:"prize_id"`
	MID       int32 `gorm:"column:m_id;primaryKey" json:"m_id"`
	CheckTime int32 `gorm:"column:check_time;not null" json:"check_time"`
}

// TableName DnfEventPrize's table name
func (*DnfEventPrize) TableName() string {
	return TableNameDnfEventPrize
}
