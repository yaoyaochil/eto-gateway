// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePswdQstionDirect = "pswd_qstion_direct"

// PswdQstionDirect mapped from table <pswd_qstion_direct>
type PswdQstionDirect struct {
	MID   int32  `gorm:"column:m_id;primaryKey" json:"m_id"`
	QText string `gorm:"column:q_text;not null" json:"q_text"`
}

// TableName PswdQstionDirect's table name
func (*PswdQstionDirect) TableName() string {
	return TableNamePswdQstionDirect
}
