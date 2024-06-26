// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameP2pnetworkStatisticDaily = "p2pnetwork_statistic_daily"

// P2pnetworkStatisticDaily mapped from table <p2pnetwork_statistic_daily>
type P2pnetworkStatisticDaily struct {
	CurDate             time.Time `gorm:"column:cur_date;primaryKey" json:"cur_date"`
	SuccessParty        float32   `gorm:"column:success_party;not null;default:0.00" json:"success_party"`
	DungeonBad          float32   `gorm:"column:dungeon_bad;not null;default:0.00" json:"dungeon_bad"`
	PvpBad              float32   `gorm:"column:pvp_bad;not null;default:0.00" json:"pvp_bad"`
	SuccessDungeonClear float32   `gorm:"column:success_dungeon_clear;not null;default:0.00" json:"success_dungeon_clear"`
	FairPvpBad          float32   `gorm:"column:fair_pvp_bad;not null;default:0.00" json:"fair_pvp_bad"`
}

// TableName P2pnetworkStatisticDaily's table name
func (*P2pnetworkStatisticDaily) TableName() string {
	return TableNameP2pnetworkStatisticDaily
}
