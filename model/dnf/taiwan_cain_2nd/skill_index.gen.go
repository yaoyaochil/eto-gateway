// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameSkillIndex = "skill_index"

// SkillIndex mapped from table <skill_index>
type SkillIndex struct {
	No        int32  `gorm:"column:no;primaryKey;autoIncrement:true" json:"no"`
	Job       int32  `gorm:"column:job;not null;default:100" json:"job"`
	SkillIdx  int32  `gorm:"column:skill_idx;not null" json:"skill_idx"`
	SkillName string `gorm:"column:skill_name;not null" json:"skill_name"`
}

// TableName SkillIndex's table name
func (*SkillIndex) TableName() string {
	return TableNameSkillIndex
}
