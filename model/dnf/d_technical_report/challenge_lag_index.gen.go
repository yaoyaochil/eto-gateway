// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameChallengeLagIndex = "challenge_lag_index"

// ChallengeLagIndex mapped from table <challenge_lag_index>
type ChallengeLagIndex struct {
	SpecID           int32     `gorm:"column:spec_id;not null" json:"spec_id"`
	OccTime          time.Time `gorm:"column:occ_time;not null" json:"occ_time"`
	ServerGroup      int32     `gorm:"column:server_group;not null" json:"server_group"`
	MinFps           int32     `gorm:"column:min_fps;not null" json:"min_fps"`
	AvgFps           int32     `gorm:"column:avg_fps;not null" json:"avg_fps"`
	MaxFps           int32     `gorm:"column:max_fps;not null" json:"max_fps"`
	WinFps           int32     `gorm:"column:win_fps;not null" json:"win_fps"`
	FullFps          int32     `gorm:"column:full_fps;not null" json:"full_fps"`
	FullWinFps       int32     `gorm:"column:full_win_fps;not null" json:"full_win_fps"`
	FullWinNosyncFps int32     `gorm:"column:full_win_nosync_fps;not null" json:"full_win_nosync_fps"`
	Frame1           int32     `gorm:"column:frame1;not null" json:"frame1"`
	Time1            float32   `gorm:"column:time1;not null;default:0.000" json:"time1"`
	Frame2           int32     `gorm:"column:frame2;not null" json:"frame2"`
	Time2            float32   `gorm:"column:time2;not null;default:0.000" json:"time2"`
	Frame3           int32     `gorm:"column:frame3;not null" json:"frame3"`
	Time3            float32   `gorm:"column:time3;not null;default:0.000" json:"time3"`
	Frame4           int32     `gorm:"column:frame4;not null" json:"frame4"`
	Time4            float32   `gorm:"column:time4;not null;default:0.000" json:"time4"`
	Frame5           int32     `gorm:"column:frame5;not null" json:"frame5"`
	Time5            float32   `gorm:"column:time5;not null;default:0.000" json:"time5"`
	Frame6           int32     `gorm:"column:frame6;not null" json:"frame6"`
	Time6            float32   `gorm:"column:time6;not null;default:0.000" json:"time6"`
	ShareRate        int32     `gorm:"column:share_rate;not null" json:"share_rate"`
}

// TableName ChallengeLagIndex's table name
func (*ChallengeLagIndex) TableName() string {
	return TableNameChallengeLagIndex
}
