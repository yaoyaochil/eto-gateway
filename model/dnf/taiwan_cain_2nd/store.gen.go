// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameStore = "store"

// Store mapped from table <store>
type Store struct {
	CharacNo int32  `gorm:"column:charac_no;primaryKey" json:"charac_no"`
	Store    []byte `gorm:"column:store;not null" json:"store"`
	UseDoll  bool   `gorm:"column:use_doll" json:"use_doll"`
}

// TableName Store's table name
func (*Store) TableName() string {
	return TableNameStore
}
