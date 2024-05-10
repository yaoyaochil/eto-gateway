package taiwan_billing

import (
	"fmt"
	"gateway/model/common/request/taiwan_billing_cera_request"
	"gateway/model/common/response"
	model "gateway/model/dnf/taiwan_billing"
	"gateway/service"
	"github.com/gin-gonic/gin"
)

type CashCeraApi struct{}

var cashCeraService = service.GroupServiceApp.TaiwanBilling.CeraService

// @Summary 取得CERA点数
// @Produce json
// @Param account query string true "账号"
// @Success 200 {object} CashCeraResponse
// @Router /taiwan_billing/cash_cera [get]
func (cera *CashCeraApi) GetCashCera(c *gin.Context) {
	var info model.CashCera
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage("参数绑定失败", c)
		return
	}
	data, err := cashCeraService.GetCera(info)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(data, c)
}

// @Summary 充值CERA点数
// @Produce json
// @Param account query string true "账号"
// @Param cera query int true "点数"
// @Success 200 {object} CashCeraResponse
// @Router /taiwan_billing/cash_cera [post]
func (cera *CashCeraApi) RechargeCashCera(c *gin.Context) {
	var info model.CashCera
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage("参数绑定失败", c)
		return
	}
	data, err := cashCeraService.RechargeCera(info)
	if err != nil {
		response.FailWithMessage("充值失敗", c)
		return
	}
	response.OkWithData(data, c)
}

// @Summary 扣减CERA点数
// @Produce json
// @Param account query string true "账号"
// @Param cera query int true "点数"
// @Success 200 {object} CashCeraResponse
// @Router /taiwan_billing/cash_cera [delete]
func (cera *CashCeraApi) DeductCashCera(c *gin.Context) {
	var info model.CashCera
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage("參數綁定失敗", c)
		return
	}
	data, err := cashCeraService.DeductCera(info)
	if err != nil {
		if err.Error() == "点券不足" {
			response.FailWithMessage("点券不足", c)
			return
		}
		response.FailWithMessage(fmt.Sprintf("扣减失败，%v", err), c)
	}
	response.OkWithData(data, c)
}

// @Summary 分页获取账号点券
// @Produce json
// @Param account query string false "账号"
// @Param pageSize query int true "每页数量"
// @Param page query int true "页码"
// @Success 200 {object} CashCeraListResponse
// @Router /taiwan_billing/cash_cera/list [get]
func (cera *CashCeraApi) GetCashCeraList(c *gin.Context) {
	var info taiwan_billing_cera_request.CashCeraRequest
	if err := c.ShouldBindJSON(&info); err != nil {
		response.FailWithMessage("参数绑定失败", c)
		return
	}
	list, total, err := cashCeraService.GetCeraList(info)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:  list,
		Total: total,
	}, "获取成功", c)
}
