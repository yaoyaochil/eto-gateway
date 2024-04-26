// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameChurnSystemManager = "churn_system_manager"

// ChurnSystemManager mapped from table <churn_system_manager>
type ChurnSystemManager struct {
	No            int32     `gorm:"column:no;primaryKey;autoIncrement:true" json:"no"`
	WeekdayVarA   int32     `gorm:"column:weekday_var_a;not null" json:"weekday_var_a"`
	WeekdayVarB   int32     `gorm:"column:weekday_var_b;not null" json:"weekday_var_b"`
	WeekdayVarC   int32     `gorm:"column:weekday_var_c;not null" json:"weekday_var_c"`
	WeekendVarX   int32     `gorm:"column:weekend_var_x;not null" json:"weekend_var_x"`
	WeekendVarY   int32     `gorm:"column:weekend_var_y;not null" json:"weekend_var_y"`
	WeekendVarZ   int32     `gorm:"column:weekend_var_z;not null" json:"weekend_var_z"`
	NextRewardDay int32     `gorm:"column:next_reward_day;not null" json:"next_reward_day"`
	AdminID       int32     `gorm:"column:admin_id;not null" json:"admin_id"`
	RegTime       time.Time `gorm:"column:reg_time;not null" json:"reg_time"`
	StateFlag     int32     `gorm:"column:state_flag;not null" json:"state_flag"`
}

// TableName ChurnSystemManager's table name
func (*ChurnSystemManager) TableName() string {
	return TableNameChurnSystemManager
}