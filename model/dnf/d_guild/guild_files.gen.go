// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGuildFile = "guild_files"

// GuildFile mapped from table <guild_files>
type GuildFile struct {
	Gno          int32  `gorm:"column:gno;primaryKey" json:"gno"`
	GfNo         int32  `gorm:"column:gf_no;primaryKey" json:"gf_no"`
	FileServer   string `gorm:"column:file_server;not null" json:"file_server"`
	FileLocation string `gorm:"column:file_location;not null" json:"file_location"`
}

// TableName GuildFile's table name
func (*GuildFile) TableName() string {
	return TableNameGuildFile
}