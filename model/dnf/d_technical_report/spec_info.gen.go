// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSpecInfo = "spec_info"

// SpecInfo mapped from table <spec_info>
type SpecInfo struct {
	UID        int32  `gorm:"column:uid;primaryKey;autoIncrement:true" json:"uid"`
	VendorID   int32  `gorm:"column:vendor_id;not null" json:"vendor_id"`
	DeviceID   int32  `gorm:"column:device_id;not null" json:"device_id"`
	VendorName string `gorm:"column:vendor_name;not null" json:"vendor_name"`
	DeviceName string `gorm:"column:device_name;not null" json:"device_name"`
}

// TableName SpecInfo's table name
func (*SpecInfo) TableName() string {
	return TableNameSpecInfo
}
