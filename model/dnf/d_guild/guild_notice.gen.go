// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGuildNotice = "guild_notice"

// GuildNotice mapped from table <guild_notice>
type GuildNotice struct {
	GuildID int32  `gorm:"column:guild_id;primaryKey" json:"guild_id"`
	Notice  string `gorm:"column:notice;not null" json:"notice"`
	AccDate int32  `gorm:"column:acc_date;not null" json:"acc_date"`
}

// TableName GuildNotice's table name
func (*GuildNotice) TableName() string {
	return TableNameGuildNotice
}