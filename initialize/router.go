package initialize

import (
	"gateway/middleware"
	"gateway/model/common/response"
	"gateway/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	// 自定义404
	Router.NoRoute(func(c *gin.Context) {
		response.NotFound(c)
	})

	//systemRouter := router.RouterGroupApp.System
	systemRouter := router.RouterGroupApp.System
	d_taiwan_account_router := router.RouterGroupApp.DTaiwanAccount

	// 方便统一添加路由组前缀 多服务器上线使用
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		// 方便统一添加路由组前缀 多服务器上线使用
		systemRouter.SystemBaseRouter.InitSystemBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.SysJwtRouter.InitJwtRouter(PublicGroup)            // 注册jwt相关路由

		// 下面是游戏相关的路由
		d_taiwan_account_router.AccountRouter.InitDTaiwanAccountRouter(PublicGroup) // 注册d_taiwan账号管理
	}
	{
		// 需要鉴权的路由
		systemRouter.SysUserRouter.InitSysUserRouter(PrivateGroup) // 注册用户路由
	}

	return Router
}
