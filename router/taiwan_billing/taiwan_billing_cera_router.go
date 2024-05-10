package taiwan_billing

import (
	v1 "gateway/api/v1"
	"github.com/gin-gonic/gin"
)

type CeraRouter struct{}

func (c *CeraRouter) InitCeraRouter(router *gin.RouterGroup) {
	var ceraApi = v1.ApiGroupApp.TaiwanBillingApiGroup
	routerGroup := router.Group("taiwan_billing")
	{
		routerGroup.POST("getCeraList", ceraApi.GetCashCeraList)   // 获取账号点券列表
		routerGroup.POST("getCera", ceraApi.GetCashCera)           // 获取账号点券
		routerGroup.POST("rechargeCera", ceraApi.RechargeCashCera) // 充值点券
		routerGroup.POST("deductCera", ceraApi.DeductCashCera)     // 扣减点券
	}

}
