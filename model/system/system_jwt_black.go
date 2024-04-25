package system

import "gateway/global"

type JwtBlacklist struct {
	global.GateWayMODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
