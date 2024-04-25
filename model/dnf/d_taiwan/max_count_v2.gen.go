// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMaxCountV2 = "max_count_v2"

// MaxCountV2 mapped from table <max_count_v2>
type MaxCountV2 struct {
	ServerInfo               int32     `gorm:"column:server_info;not null" json:"server_info"`
	NumOccupationsCharscreen int32     `gorm:"column:num_occupations_charscreen;not null" json:"num_occupations_charscreen"`
	NumOccupationsSeriaroom  int32     `gorm:"column:num_occupations_seriaroom;not null" json:"num_occupations_seriaroom"`
	NumLoginPerMin           int32     `gorm:"column:num_login_per_min;not null" json:"num_login_per_min"`
	NumLogoutPerMin          int32     `gorm:"column:num_logout_per_min;not null" json:"num_logout_per_min"`
	McDate                   time.Time `gorm:"column:mc_date;not null" json:"mc_date"`
}

// TableName MaxCountV2's table name
func (*MaxCountV2) TableName() string {
	return TableNameMaxCountV2
}
