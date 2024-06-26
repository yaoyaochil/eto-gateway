// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameEventTowerEntry = "event_tower_entry"

// EventTowerEntry mapped from table <event_tower_entry>
type EventTowerEntry struct {
	MID        int32 `gorm:"column:m_id;primaryKey" json:"m_id"`
	OccDate    int32 `gorm:"column:occ_date;not null" json:"occ_date"`
	OccCheck   int32 `gorm:"column:occ_check;not null" json:"occ_check"`
	ServerID   int32 `gorm:"column:server_id;not null" json:"server_id"`
	CharacNo   int32 `gorm:"column:charac_no;not null" json:"charac_no"`
	Item1No    int32 `gorm:"column:item1_no;not null" json:"item1_no"`
	Item1Check int32 `gorm:"column:item1_check;not null" json:"item1_check"`
	Item2No    int32 `gorm:"column:item2_no;not null" json:"item2_no"`
	Item2Check int32 `gorm:"column:item2_check;not null" json:"item2_check"`
	Item3No    int32 `gorm:"column:item3_no;not null" json:"item3_no"`
	Item3Check int32 `gorm:"column:item3_check;not null" json:"item3_check"`
}

// TableName EventTowerEntry's table name
func (*EventTowerEntry) TableName() string {
	return TableNameEventTowerEntry
}
