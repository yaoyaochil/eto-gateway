package system

import (
	v1 "gateway/api/v1"

	"github.com/gin-gonic/gin"
)

type SysJwtRouter struct{}

func (s *SysJwtRouter) InitJwtRouter(Router *gin.RouterGroup) {
	sysJwtRouter := Router.Group("jwt")
	jwtApi := v1.ApiGroupApp.SystemApiGroup.SysJwtApi
	{
		sysJwtRouter.POST("jwtInBlacklist", jwtApi.JwtInBlacklist) // jwt作废
	}
}
