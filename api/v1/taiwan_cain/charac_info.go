package taiwan_cain

import (
	"gateway/model/common/request/taiwan_cain_request"
	"gateway/model/common/response"
	model "gateway/model/dnf/taiwan_cain"
	"gateway/service"
	"github.com/gin-gonic/gin"
)

type CharacInfoApi struct {
}

var characInfoServer = service.GroupServiceApp.TaiwanCain

// GetCharacInfo 获取角色信息
func (charac *CharacInfoApi) GetCharacInfo(c *gin.Context) {
	var info model.CharacInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := characInfoServer.TaiwanCainCharacterInfoService.GetCharacterInfo(info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(data, c)
}

// GetCharacInfoList 获取角色列表
func (charac *CharacInfoApi) GetCharacInfoList(c *gin.Context) {
	var info taiwan_cain_request.TaiwanCainCharacterInfoRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, total, err := characInfoServer.TaiwanCainCharacterInfoService.GetCharacterInfoList(info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     data,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "获取成功", c)
}
