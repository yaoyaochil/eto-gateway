// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSlangList = "slang_list"

// SlangList mapped from table <slang_list>
type SlangList struct {
	Slang string `gorm:"column:slang;primaryKey" json:"slang"`
}

// TableName SlangList's table name
func (*SlangList) TableName() string {
	return TableNameSlangList
}
