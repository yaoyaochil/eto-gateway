package system

import (
	"gateway/global"

	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	global.GateWayMODEL
	UUID      uuid.UUID `json:"uuid,omitempty" gorm:"comment:用户UUID"`
	Username  string    `json:"username" gorm:"comment:用户登录名;index;unique;"`
	Password  string    `json:"password,omitempty" gorm:"comment:用户登录密码"`
	NickName  string    `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	HeaderImg string    `json:"headerImg" gorm:"default:https://huaduogiegie.oss-cn-shanghai.aliyuncs.com/picgo/202305281629232.png;comment:用户头像"`
}
