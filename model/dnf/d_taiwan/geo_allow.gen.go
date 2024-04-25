// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGeoAllow = "geo_allow"

// GeoAllow mapped from table <geo_allow>
type GeoAllow struct {
	AllowIP    string    `gorm:"column:allow_ip;primaryKey" json:"allow_ip"`
	AllowCCode string    `gorm:"column:allow_c_code;not null" json:"allow_c_code"`
	AllowDate  time.Time `gorm:"column:allow_date;not null;default:CURRENT_TIMESTAMP" json:"allow_date"`
}

// TableName GeoAllow's table name
func (*GeoAllow) TableName() string {
	return TableNameGeoAllow
}
