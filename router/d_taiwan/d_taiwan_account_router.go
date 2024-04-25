package d_taiwan

import (
	v1 "gateway/api/v1"
	"github.com/gin-gonic/gin"
)

type AccountRouter struct{}

// InitDTaiwanAccountRouter 注册d_taiwan账号管理
func (s *AccountRouter) InitDTaiwanAccountRouter(RouterGroup *gin.RouterGroup) {
	var dTaiwanApi = v1.ApiGroupApp.DTaiwanApiGroup.AccountApi
	Router := RouterGroup.Group("/d_taiwan")
	{
		Router.POST("getAccount", dTaiwanApi.GetAccount)
		Router.POST("dnfLogin", dTaiwanApi.DnfLogin)
	}
}
