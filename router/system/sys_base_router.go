package system

import (
	v1 "gateway/api/v1"

	"github.com/gin-gonic/gin"
)

type SystemBaseRouter struct {
}

func (s *SystemBaseRouter) InitSystemBaseRouter(router *gin.RouterGroup) {
	Router := router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup
	{
		Router.POST("login", baseApi.Login)
	}
}
