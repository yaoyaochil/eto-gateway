// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGuildJoinList = "guild_join_list"

// GuildJoinList mapped from table <guild_join_list>
type GuildJoinList struct {
	GuildID     int32     `gorm:"column:guild_id;primaryKey" json:"guild_id"`
	CharacNo    int32     `gorm:"column:charac_no;primaryKey" json:"charac_no"`
	ServerGroup int32     `gorm:"column:server_group;not null" json:"server_group"`
	MID         int32     `gorm:"column:m_id;not null" json:"m_id"`
	BornYear    string    `gorm:"column:born_year;not null" json:"born_year"`
	Memo        string    `gorm:"column:memo;not null" json:"memo"`
	OccTime     time.Time `gorm:"column:occ_time" json:"occ_time"`
}

// TableName GuildJoinList's table name
func (*GuildJoinList) TableName() string {
	return TableNameGuildJoinList
}
