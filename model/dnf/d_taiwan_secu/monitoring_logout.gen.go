// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMonitoringLogout = "monitoring_logout"

// MonitoringLogout mapped from table <monitoring_logout>
type MonitoringLogout struct {
	No         int32 `gorm:"column:no;primaryKey;autoIncrement:true" json:"no"`
	MID        int32 `gorm:"column:m_id;not null" json:"m_id"`
	LogoutTime int32 `gorm:"column:logout_time;not null" json:"logout_time"`
	LogoutIP   int32 `gorm:"column:logout_ip;not null" json:"logout_ip"`
	OtpDelType int32 `gorm:"column:otp_del_type;not null" json:"otp_del_type"`
}

// TableName MonitoringLogout's table name
func (*MonitoringLogout) TableName() string {
	return TableNameMonitoringLogout
}
