// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUserItem = "user_items"

// UserItem mapped from table <user_items>
type UserItem struct {
	UIID            int32     `gorm:"column:ui_id;primaryKey;autoIncrement:true" json:"ui_id"`
	CharacNo        int32     `gorm:"column:charac_no;not null" json:"charac_no"`
	Slot            int32     `gorm:"column:slot;not null" json:"slot"`
	ItID            int32     `gorm:"column:it_id;not null" json:"it_id"`
	ExpireDate      time.Time `gorm:"column:expire_date;not null" json:"expire_date"`
	ObtainFrom      int32     `gorm:"column:obtain_from" json:"obtain_from"`
	RegDate         time.Time `gorm:"column:reg_date;not null" json:"reg_date"`
	IpgAgencyNo     string    `gorm:"column:ipg_agency_no;not null" json:"ipg_agency_no"`
	AbilityNo       int32     `gorm:"column:ability_no;not null" json:"ability_no"`
	Stat            int32     `gorm:"column:stat;not null" json:"stat"`
	ClearAvatarID   int32     `gorm:"column:clear_avatar_id;not null" json:"clear_avatar_id"`
	JewelSocket     []byte    `gorm:"column:jewel_socket;not null" json:"jewel_socket"`
	ItemLockKey     int32     `gorm:"column:item_lock_key;not null" json:"item_lock_key"`
	ToIpgAgencyNo   string    `gorm:"column:to_ipg_agency_no;not null" json:"to_ipg_agency_no"`
	MTime           time.Time `gorm:"column:m_time;not null" json:"m_time"`
	HiddenOption    int32     `gorm:"column:hidden_option;not null" json:"hidden_option"`
	EmblemEndurance int32     `gorm:"column:emblem_endurance;not null" json:"emblem_endurance"`
	Color1          int32     `gorm:"column:color1" json:"color1"`
	Color2          int32     `gorm:"column:color2" json:"color2"`
	TradeRestrict   int32     `gorm:"column:trade_restrict" json:"trade_restrict"`
}

// TableName UserItem's table name
func (*UserItem) TableName() string {
	return TableNameUserItem
}
