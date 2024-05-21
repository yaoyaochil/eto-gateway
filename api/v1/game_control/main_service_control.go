package gamecontrol

import (
	"gateway/model/common/response"
	"gateway/service"

	"github.com/gin-gonic/gin"
)

type MainServiceControlApi struct{}

var main_service = service.GroupServiceApp.GameControl.ServiceControl

func (msc *MainServiceControlApi) StartMainServices(c *gin.Context) {
	if err := main_service.RunMainService(); err != nil {
		response.FailWithMessage("启动失败", c)
		return
	}
	response.OkWithMessage("启动成功", c)
}

func (msc *MainServiceControlApi) StopMainServices(c *gin.Context) {
	if err := main_service.StopMainService(); err != nil {
		response.FailWithMessage("停止失败", c)
		return
	}
	response.OkWithMessage("停止成功", c)
}

func (msc *MainServiceControlApi) GetServiceStatus(c *gin.Context) {
	status := main_service.CheckProcess()

	if len(status) == 0 {
		response.FailWithMessage("服务未启动", c)
		return
	}

	response.OkWithData(status, c)
}
