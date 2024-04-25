// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTmeCharac = "tme_charac"

// TmeCharac mapped from table <tme_charac>
type TmeCharac struct {
	MID               int32     `gorm:"column:m_id;not null" json:"m_id"`
	CharacNo          int32     `gorm:"column:charac_no;not null" json:"charac_no"`
	CharacName        string    `gorm:"column:charac_name;not null" json:"charac_name"`
	Village           int32     `gorm:"column:village;not null;default:1" json:"village"`
	Job               int32     `gorm:"column:job;not null" json:"job"`
	Lev               int32     `gorm:"column:lev;not null;default:1" json:"lev"`
	Exp               int32     `gorm:"column:exp;not null" json:"exp"`
	GrowType          int32     `gorm:"column:grow_type;not null" json:"grow_type"`
	HP                int32     `gorm:"column:HP;not null" json:"HP"`
	MaxHP             int32     `gorm:"column:maxHP;not null" json:"maxHP"`
	MaxMP             int32     `gorm:"column:maxMP;not null" json:"maxMP"`
	PhyAttack         int32     `gorm:"column:phy_attack;not null" json:"phy_attack"`
	PhyDefense        int32     `gorm:"column:phy_defense;not null" json:"phy_defense"`
	MagAttack         int32     `gorm:"column:mag_attack;not null" json:"mag_attack"`
	MagDefense        int32     `gorm:"column:mag_defense;not null" json:"mag_defense"`
	ElementResist     []byte    `gorm:"column:element_resist;not null" json:"element_resist"`
	SpecProperty      []byte    `gorm:"column:spec_property;not null" json:"spec_property"`
	InvenWeight       int32     `gorm:"column:inven_weight;not null" json:"inven_weight"`
	HpRegen           int32     `gorm:"column:hp_regen;not null" json:"hp_regen"`
	MpRegen           int32     `gorm:"column:mp_regen;not null" json:"mp_regen"`
	MoveSpeed         int32     `gorm:"column:move_speed;not null" json:"move_speed"`
	AttackSpeed       int32     `gorm:"column:attack_speed;not null" json:"attack_speed"`
	CastSpeed         int32     `gorm:"column:cast_speed;not null" json:"cast_speed"`
	HitRecovery       int32     `gorm:"column:hit_recovery;not null" json:"hit_recovery"`
	Jump              int32     `gorm:"column:jump;not null" json:"jump"`
	CharacWeight      int32     `gorm:"column:charac_weight;not null" json:"charac_weight"`
	Fatigue           int32     `gorm:"column:fatigue;not null" json:"fatigue"`
	MaxFatigue        int32     `gorm:"column:max_fatigue;not null;default:70" json:"max_fatigue"`
	PremiumFatigue    int32     `gorm:"column:premium_fatigue;not null" json:"premium_fatigue"`
	MaxPremiumFatigue int32     `gorm:"column:max_premium_fatigue;not null" json:"max_premium_fatigue"`
	CreateTime        time.Time `gorm:"column:create_time;not null" json:"create_time"`
	LastPlayTime      time.Time `gorm:"column:last_play_time;not null" json:"last_play_time"`
	DungeonClearPoint int32     `gorm:"column:dungeon_clear_point;not null" json:"dungeon_clear_point"`
	DeleteTime        time.Time `gorm:"column:delete_time;not null" json:"delete_time"`
	DeleteFlag        int32     `gorm:"column:delete_flag;not null" json:"delete_flag"`
	GuildID           int32     `gorm:"column:guild_id;not null" json:"guild_id"`
	GuildRight        int32     `gorm:"column:guild_right;not null" json:"guild_right"`
	MemberFlag        int32     `gorm:"column:member_flag;not null" json:"member_flag"`
}

// TableName TmeCharac's table name
func (*TmeCharac) TableName() string {
	return TableNameTmeCharac
}
