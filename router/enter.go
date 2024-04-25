package router

import (
	"gateway/router/d_taiwan"
	"gateway/router/system"
)

type RouterGroup struct {
	System         system.RouterGroup
	DTaiwanAccount d_taiwan.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
