package v1

import (
	"gateway/api/v1/d_taiwan"
	"gateway/api/v1/system"
	"gateway/api/v1/taiwan_billing"
	"gateway/api/v1/taiwan_cain"
	"gateway/api/v1/taiwan_cain_2nd"
	"gateway/api/v1/websocket"
)

type ApiGroup struct {
	SystemApiGroup        system.ApiGroup          // system 模块
	DTaiwanApiGroup       d_taiwan.ApiGroup        // d_taiwan 模块
	TaiwanBillingApiGroup taiwan_billing.ApiGroup  // taiwan_billing 模块
	TaiwanCainApiGroup    taiwan_cain.ApiGroup     // taiwan_cain 模块
	TaiwanCain2NDApiGroup taiwan_cain_2nd.ApiGroup // taiwan_cain_2nd 模块
	WebsocketApiGroup     websocket.ApiGroup       // websocket 模块
}

var ApiGroupApp = new(ApiGroup)
