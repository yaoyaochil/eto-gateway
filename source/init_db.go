package source

import (
	"gateway/global"
	"gateway/model/system"
	"gateway/utils"

	uuid "github.com/satori/go.uuid"
)

func InitSystemDB() {
	db, exists := global.GatewayDBs["system"]
	if !exists || db == nil {
		global.GatewayLog.Error("GormMysqlSystem DB is nil")
		return
	}

	var userList system.SysUser
	db.Where("username = ?", "admin").First(&userList)
	if userList.ID != 0 {
		global.GatewayLog.Info("系统用户admin已存在 本次无需初始化用户!")
		return
	}

	// Init system db
	user := system.SysUser{
		UUID:      uuid.NewV4(),
		Username:  "admin",
		Password:  utils.BcryptHash("123456"),
		NickName:  "系统管理员",
		HeaderImg: "https://img.imdodo.com/openapitest/upload/cdn/4A301072DEC6B6A49050E5B294CD7983_1707877949352.png",
	}

	db.Create(&user)
}
