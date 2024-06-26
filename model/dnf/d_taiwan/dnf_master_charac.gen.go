// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameDnfMasterCharac = "dnf_master_charac"

// DnfMasterCharac mapped from table <dnf_master_charac>
type DnfMasterCharac struct {
	MID        int32  `gorm:"column:m_id;primaryKey" json:"m_id"`
	GlobalType int32  `gorm:"column:global_type;primaryKey" json:"global_type"`
	ServerID   int32  `gorm:"column:server_id;primaryKey" json:"server_id"`
	CharacNo   int32  `gorm:"column:charac_no;not null" json:"charac_no"`
	CharacName string `gorm:"column:charac_name;not null" json:"charac_name"`
	Job        int32  `gorm:"column:job;not null" json:"job"`
	Lev        int32  `gorm:"column:lev;not null" json:"lev"`
}

// TableName DnfMasterCharac's table name
func (*DnfMasterCharac) TableName() string {
	return TableNameDnfMasterCharac
}
