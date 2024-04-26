// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGuildStatMonth = "guild_stat_month"

// GuildStatMonth mapped from table <guild_stat_month>
type GuildStatMonth struct {
	OccDate          time.Time `gorm:"column:occ_date;primaryKey" json:"occ_date"`
	Lev              int32     `gorm:"column:lev;primaryKey" json:"lev"`
	ServerID         int32     `gorm:"column:server_id;primaryKey" json:"server_id"`
	AvgGuildPoint    int32     `gorm:"column:avg_guild_point;not null" json:"avg_guild_point"`
	AvgGuildPointAcc int32     `gorm:"column:avg_guild_point_acc;not null" json:"avg_guild_point_acc"`
}

// TableName GuildStatMonth's table name
func (*GuildStatMonth) TableName() string {
	return TableNameGuildStatMonth
}