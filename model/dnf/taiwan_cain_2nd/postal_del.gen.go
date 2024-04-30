// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePostalDel = "postal_del"

// PostalDel mapped from table <postal_del>
type PostalDel struct {
	Sdate            time.Time `gorm:"column:sdate;primaryKey" json:"sdate"`
	PostalID         int32     `gorm:"column:postal_id;primaryKey" json:"postal_id"`
	OccTime          time.Time `gorm:"column:occ_time;not null" json:"occ_time"`
	SendCharacNo     int32     `gorm:"column:send_charac_no;not null" json:"send_charac_no"`
	SendCharacName   string    `gorm:"column:send_charac_name;not null" json:"send_charac_name"`
	ReceiveCharacNo  int32     `gorm:"column:receive_charac_no;not null" json:"receive_charac_no"`
	ItemID           int32     `gorm:"column:item_id;not null" json:"item_id"`
	AddInfo          int32     `gorm:"column:add_info;not null" json:"add_info"`
	Endurance        int32     `gorm:"column:endurance;not null" json:"endurance"`
	Upgrade          int32     `gorm:"column:upgrade;not null" json:"upgrade"`
	AmplifyOption    int32     `gorm:"column:amplify_option;not null" json:"amplify_option"`
	AmplifyValue     int32     `gorm:"column:amplify_value;not null" json:"amplify_value"`
	Gold             int32     `gorm:"column:gold;not null" json:"gold"`
	ReceiveTime      time.Time `gorm:"column:receive_time;not null" json:"receive_time"`
	DeleteFlag       int32     `gorm:"column:delete_flag;not null" json:"delete_flag"`
	AvataFlag        int32     `gorm:"column:avata_flag;not null" json:"avata_flag"`
	UnlimitFlag      int32     `gorm:"column:unlimit_flag;not null" json:"unlimit_flag"`
	SealFlag         int32     `gorm:"column:seal_flag;not null" json:"seal_flag"`
	CreatureFlag     int32     `gorm:"column:creature_flag;not null" json:"creature_flag"`
	Postal           int32     `gorm:"column:postal;not null" json:"postal"`
	LetterID         int32     `gorm:"column:letter_id;not null" json:"letter_id"`
	ExtendInfo       int32     `gorm:"column:extend_info;not null" json:"extend_info"`
	IpgDbID          int32     `gorm:"column:ipg_db_id;not null" json:"ipg_db_id"`
	IpgTransactionID int32     `gorm:"column:ipg_transaction_id;not null" json:"ipg_transaction_id"`
	IpgNexonID       string    `gorm:"column:ipg_nexon_id;not null" json:"ipg_nexon_id"`
	AuctionID        int64     `gorm:"column:auction_id;not null" json:"auction_id"`
	RandomOption     []byte    `gorm:"column:random_option;not null" json:"random_option"`
	SeperateUpgrade  int32     `gorm:"column:seperate_upgrade;not null" json:"seperate_upgrade"`
	Type             int32     `gorm:"column:type;not null" json:"type"`
	ItemGUID         []byte    `gorm:"column:item_guid;not null" json:"item_guid"`
}

// TableName PostalDel's table name
func (*PostalDel) TableName() string {
	return TableNamePostalDel
}
