// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCharacItemLockInfo = "charac_item_lock_info"

// CharacItemLockInfo mapped from table <charac_item_lock_info>
type CharacItemLockInfo struct {
	CharacNo     int32  `gorm:"column:charac_no;primaryKey" json:"charac_no"`
	ItemLockInfo []byte `gorm:"column:item_lock_info" json:"item_lock_info"`
}

// TableName CharacItemLockInfo's table name
func (*CharacItemLockInfo) TableName() string {
	return TableNameCharacItemLockInfo
}