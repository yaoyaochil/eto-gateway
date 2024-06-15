package base

import (
	v1 "gateway/api/v1"

	"github.com/gin-gonic/gin"
)

type DTaiwanAccountRouter struct{}

func (a *DTaiwanAccountRouter) InitDTaiwanAccountRouter(router *gin.RouterGroup) {
	var accountApi = v1.ApiGroupApp.DTaiwanApiGroup.AccountApi
	router = router.Group("d_taiwan")
	{
		router.POST("register", accountApi.Register)
		router.POST("dnfLogin", accountApi.DnfLogin)
		router.POST("loginProtobuf", accountApi.DnfLoginProtobuf)
	}

}
