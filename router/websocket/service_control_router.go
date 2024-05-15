package websocket

import (
	v1 "gateway/api/v1"

	"github.com/gin-gonic/gin"
)

type RouterServiceControl struct{}

func (r *RouterServiceControl) InitServiceControlRouter(router *gin.RouterGroup) {
	var api = v1.ApiGroupApp.WebsocketApiGroup.ServiceControlApi
	{
		router.GET("/service_control", api.UpgradeServiceControlApi)
	}
}
