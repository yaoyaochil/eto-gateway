// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCharacAdvanceAltarItemDesc = "charac_advance_altar_item_desc"

// CharacAdvanceAltarItemDesc mapped from table <charac_advance_altar_item_desc>
type CharacAdvanceAltarItemDesc struct {
	RidableID int32  `gorm:"column:ridable_id;primaryKey;comment:ｺｯｽﾅｸ?ｺﾅﾍ ｾﾆﾀﾌｵ" json:"ridable_id"`                // ｺｯｽﾅｸ?ｺﾅﾍ ｾﾆﾀﾌｵ
	ItemType  int32  `gorm:"column:item_type;primaryKey;comment:ｾﾆﾀﾌﾅﾛﾅｸﾀﾔ 0:ﾀｯｴﾖ, 1:ｽｺﾅｳ, 2:ﾅｸｿ" json:"item_type"` // ｾﾆﾀﾌﾅﾛﾅｸﾀﾔ 0:ﾀｯｴﾖ, 1:ｽｺﾅｳ, 2:ﾅｸｿ
	ItemID    int32  `gorm:"column:item_id;primaryKey;comment:ｾﾆﾀﾌﾅﾛｾﾆﾀﾌｵ" json:"item_id"`                          // ｾﾆﾀﾌﾅﾛｾﾆﾀﾌｵ
	ItemDesc  []byte `gorm:"column:item_desc;not null;comment:ｾﾆﾀﾌﾅﾛｼｳｸ" json:"item_desc"`                          // ｾﾆﾀﾌﾅﾛｼｳｸ
}

// TableName CharacAdvanceAltarItemDesc's table name
func (*CharacAdvanceAltarItemDesc) TableName() string {
	return TableNameCharacAdvanceAltarItemDesc
}
