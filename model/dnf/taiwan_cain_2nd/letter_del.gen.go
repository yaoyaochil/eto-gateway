// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameLetterDel = "letter_del"

// LetterDel mapped from table <letter_del>
type LetterDel struct {
	Sdate          time.Time `gorm:"column:sdate;primaryKey" json:"sdate"`
	LetterID       int32     `gorm:"column:letter_id;primaryKey" json:"letter_id"`
	CharacNo       int32     `gorm:"column:charac_no;not null" json:"charac_no"`
	SendCharacNo   int32     `gorm:"column:send_charac_no;not null" json:"send_charac_no"`
	SendCharacName string    `gorm:"column:send_charac_name;not null" json:"send_charac_name"`
	LetterText     string    `gorm:"column:letter_text;not null" json:"letter_text"`
	RegDate        time.Time `gorm:"column:reg_date;not null" json:"reg_date"`
	Stat           int32     `gorm:"column:stat;not null" json:"stat"`
}

// TableName LetterDel's table name
func (*LetterDel) TableName() string {
	return TableNameLetterDel
}