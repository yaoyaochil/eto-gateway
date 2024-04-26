package router

import (
	"gateway/router/base"
	"gateway/router/d_taiwan"
	"gateway/router/system"
)

type RouterGroup struct {
	System         system.RouterGroup
	DTaiwanAccount d_taiwan.RouterGroup
	Base           base.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
