// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGeoCountryCode = "geo_country_code"

// GeoCountryCode mapped from table <geo_country_code>
type GeoCountryCode struct {
	CodeNo        int32  `gorm:"column:code_no;primaryKey" json:"code_no"`
	CountryCodeA2 string `gorm:"column:country_code_a2;not null" json:"country_code_a2"`
	CountryCodeA3 string `gorm:"column:country_code_a3;not null" json:"country_code_a3"`
	Country       string `gorm:"column:country;not null" json:"country"`
}

// TableName GeoCountryCode's table name
func (*GeoCountryCode) TableName() string {
	return TableNameGeoCountryCode
}
