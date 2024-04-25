package v1

import (
	"gateway/api/v1/d_taiwan"
	"gateway/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup   // system 模块
	DTaiwanApiGroup d_taiwan.ApiGroup // d_taiwan 模块
}

var ApiGroupApp = new(ApiGroup)
