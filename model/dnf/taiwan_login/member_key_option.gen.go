// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMemberKeyOption = "member_key_option"

// MemberKeyOption mapped from table <member_key_option>
type MemberKeyOption struct {
	MID       int64  `gorm:"column:m_id;primaryKey" json:"m_id"`
	KeyType   int32  `gorm:"column:key_type;primaryKey" json:"key_type"`
	KeyOption []byte `gorm:"column:key_option;not null" json:"key_option"`
}

// TableName MemberKeyOption's table name
func (*MemberKeyOption) TableName() string {
	return TableNameMemberKeyOption
}