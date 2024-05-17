package gamecontrol

import (
	v1 "gateway/api/v1"

	"github.com/gin-gonic/gin"
)

type MainServiceControlRouter struct{}

func (msc *MainServiceControlRouter) InitMainServiceControlRouter(router *gin.RouterGroup) {
	mainServiceControl := router.Group("main_service_control")
	var api = v1.ApiGroupApp.GameControlApiGroup.MainServiceControlApi
	{
		mainServiceControl.GET("start", api.StartMainServices)
		mainServiceControl.GET("stop", api.StopMainServices)
		mainServiceControl.GET("status", api.GetServiceStatus)
	}
}
