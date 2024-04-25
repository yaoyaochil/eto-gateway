package main

import (
	"gateway/core"
	"gateway/global"
	"gateway/initialize"
	"gateway/source"
)

func main() {
	global.GatewayVp = core.Viper() // 初始化Viper
	initialize.OtherInit()          // 初始化其他
	global.GatewayLog = core.Zap()  // 初始化zap日志库
	global.GatewayDBs = initialize.Gorm()
	for _, db := range global.GatewayDBs {
		if db != nil {
			initialize.RegisterDBTables(db) // 初始化表
			sqlDB, _ := db.DB()
			defer sqlDB.Close()
		}
	}

	source.InitSystemDB() // 初始化系统用户

	core.RunWindowsServer()
}
