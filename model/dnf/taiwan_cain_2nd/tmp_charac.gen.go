// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTmpCharac = "tmp_charac"

// TmpCharac mapped from table <tmp_charac>
type TmpCharac struct {
	MID      int32 `gorm:"column:m_id;primaryKey" json:"m_id"`
	CharacNo int32 `gorm:"column:charac_no;primaryKey" json:"charac_no"`
}

// TableName TmpCharac's table name
func (*TmpCharac) TableName() string {
	return TableNameTmpCharac
}
