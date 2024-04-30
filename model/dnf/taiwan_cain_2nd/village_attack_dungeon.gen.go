// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameVillageAttackDungeon = "village_attack_dungeon"

// VillageAttackDungeon mapped from table <village_attack_dungeon>
type VillageAttackDungeon struct {
	OccDate        time.Time `gorm:"column:occ_date;primaryKey" json:"occ_date"`
	CharacNo       int32     `gorm:"column:charac_no;primaryKey" json:"charac_no"`
	AttackCount    int32     `gorm:"column:attack_count;not null" json:"attack_count"`
	RevengeDungeon int32     `gorm:"column:revenge_dungeon;not null" json:"revenge_dungeon"`
}

// TableName VillageAttackDungeon's table name
func (*VillageAttackDungeon) TableName() string {
	return TableNameVillageAttackDungeon
}
