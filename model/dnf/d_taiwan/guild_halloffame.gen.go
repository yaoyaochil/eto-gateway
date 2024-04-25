// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGuildHalloffame = "guild_halloffame"

// GuildHalloffame mapped from table <guild_halloffame>
type GuildHalloffame struct {
	FameID    int32  `gorm:"column:fame_id;primaryKey" json:"fame_id"`
	ServerID  int32  `gorm:"column:server_id;primaryKey" json:"server_id"`
	GuildID   int32  `gorm:"column:guild_id;not null" json:"guild_id"`
	GuildName string `gorm:"column:guild_name;not null" json:"guild_name"`
	FileURL   string `gorm:"column:file_url;not null;default:0" json:"file_url"`
	OpenFlag  int32  `gorm:"column:open_flag;not null" json:"open_flag"`
	MainFlag  int32  `gorm:"column:main_flag;not null" json:"main_flag"`
}

// TableName GuildHalloffame's table name
func (*GuildHalloffame) TableName() string {
	return TableNameGuildHalloffame
}
