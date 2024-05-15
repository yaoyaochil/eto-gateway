package router

import (
	"gateway/router/base"
	"gateway/router/d_taiwan"
	"gateway/router/system"
	"gateway/router/taiwan_billing"
	"gateway/router/taiwan_cain"
	"gateway/router/taiwan_cain_2nd"
	"gateway/router/websocket"
)

type RouterGroup struct {
	System         system.RouterGroup
	DTaiwanAccount d_taiwan.RouterGroup
	TaiwanBilling  taiwan_billing.RouterGroup
	TaiwanCain     taiwan_cain.RouterGroup
	TaiwanCain2Nd  taiwan_cain_2nd.RouterGroup
	Base           base.RouterGroup
	Websocket      websocket.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
