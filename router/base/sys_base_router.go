package base

import (
	v1 "gateway/api/v1"

	"github.com/gin-gonic/gin"
)

type SystemBaseRouter struct {
}

func (s *SystemBaseRouter) InitSystemBaseRouter(router *gin.RouterGroup) {
	baseApi := v1.ApiGroupApp.SystemApiGroup
	{
		router.POST("login", baseApi.Login)
	}
}
