package global

import (
	"time"

	"gorm.io/gorm"
)

type GateWayMODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
	// 更新操作人 关联用户表
	UpdatedBy string `json:"updatedBy" gorm:"comment:更新操作人;default:'系统自动'"`
}
