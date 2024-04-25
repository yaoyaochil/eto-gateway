package system

import (
	"gateway/global"
	"gateway/model/common/response"
	"gateway/model/system"
	"gateway/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysJwtApi struct{}

// @Tags System
// @Summary JWT鉴权
// @Security ApiKeyAuth
// @Produce  application/json
// @Param token query string true "token"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /system/jwt/jwtInBlacklist [get]
func (s *SysJwtApi) JwtInBlacklist(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	if token == "" {
		response.FailWithMessage("账户已注销 请勿重复操作", c)
		return
	}
	jwt := system.JwtBlacklist{Jwt: token}
	if err := service.GroupServiceApp.SystemServiceGroup.JwtService.JwtInBlacklist(jwt); err != nil {
		global.GatewayLog.Error("退出登陆失败!", zap.Error(err))
		response.FailWithMessage("退出登陆失败", c)
		return
	}
	response.OkWithMessage("账户退出登陆!", c)
}
