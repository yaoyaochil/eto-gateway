// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCsTable2 = "cs_table2"

// CsTable2 mapped from table <cs_table2>
type CsTable2 struct {
	AccountID string `gorm:"column:account_id;not null" json:"account_id"`
	CharacID  string `gorm:"column:charac_id;not null" json:"charac_id"`
}

// TableName CsTable2's table name
func (*CsTable2) TableName() string {
	return TableNameCsTable2
}