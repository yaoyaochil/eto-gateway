// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGuildVisit = "guild_visit"

// GuildVisit mapped from table <guild_visit>
type GuildVisit struct {
	GuildID    int32 `gorm:"column:guild_id;primaryKey" json:"guild_id"`
	ServerID   int32 `gorm:"column:server_id;not null" json:"server_id"`
	TotalVisit int32 `gorm:"column:total_visit;not null" json:"total_visit"`
	TodayVisit int32 `gorm:"column:today_visit;not null" json:"today_visit"`
}

// TableName GuildVisit's table name
func (*GuildVisit) TableName() string {
	return TableNameGuildVisit
}
