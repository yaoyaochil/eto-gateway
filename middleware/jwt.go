package middleware

import (
	"gateway/global"
	"gateway/model/common/response"
	"gateway/service"
	"gateway/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var jwtService = service.GroupServiceApp.SystemServiceGroup.JwtService

// JWTAuth JWT鉴权中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")

		// 如果token为空 则返回未登录或非法访问
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		// 如果token在黑名单中 则返回您的帐户异地登陆或令牌失效
		if jwtService.IsBlacklist(token) {
			response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
			c.Abort()
			return
		}

		// 解析token包含的信息
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		// 如果token即将过期 则刷新token
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.GatewayConf.JWT.ExpiresTime)
			claims.ExpiresAt = time.Now().Add(dr).Unix()
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))

			// 如果开启了多点登录功能，则进行token拉黑操作
			//if global.GatewayConf.System.UseMultipoint {
			//	RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
			//	if err != nil {
			//		global.PalLog.Error("get redis jwt failed", zap.Error(err))
			//	} else { // 当之前的取成功时才进行拉黑操作
			//		_ = jwtService.JwtInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
			//	}
			//	// 无论如何都要记录当前的活跃状态
			//	_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			//}
		}

		c.Set("claims", claims)
		c.Next()
	}
}
