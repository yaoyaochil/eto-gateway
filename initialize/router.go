package initialize

import (
	"gateway/middleware"
	"gateway/model/common/response"
	"gateway/router"
	"net/http"

	"github.com/gin-gonic/gin"
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
	taiwan_billing_router := router.RouterGroupApp.TaiwanBilling
	taiwan_cain_router := router.RouterGroupApp.TaiwanCain
	taiwan_cain_2nd_router := router.RouterGroupApp.TaiwanCain2Nd
	base_router := router.RouterGroupApp.Base
	websocket_router := router.RouterGroupApp.Websocket
	gamecontrol_router := router.RouterGroupApp.GameControl

	// 方便统一添加路由组前缀 多服务器上线使用
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	PublicGroup := Router.Group("")

	BasePublicGroup := Router.Group("base") // 基础功能路由 不做鉴权
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		// 方便统一添加路由组前缀 多服务器上线使用
		base_router.SystemBaseRouter.InitSystemBaseRouter(BasePublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.SysJwtRouter.InitJwtRouter(PublicGroup)               // 注册jwt相关路由

		// websocket类型路由
		websocket_router.RouterServiceControl.InitServiceControlRouter(PublicGroup) // 注册websocket相关路由

		// 下面是游戏相关的路由
		base_router.DTaiwanAccountRouter.InitDTaiwanAccountRouter(BasePublicGroup)            // 注册游戏登陆,注册路由
		d_taiwan_account_router.AccountRouter.InitDTaiwanAccountRouter(PublicGroup)           // 注册d_taiwan账号管理
		taiwan_billing_router.CeraRouter.InitCeraRouter(PublicGroup)                          // 注册游戏计费相关路由
		taiwan_cain_router.CharacInfoRouter.InitCharacInfoRouter(PublicGroup)                 // 注册游戏角色相关路由
		taiwan_cain_2nd_router.EmailRouter.InitEmailRouter(PublicGroup)                       // 注册游戏第二版相关路由 邮件
		gamecontrol_router.MainServiceControlRouter.InitMainServiceControlRouter(PublicGroup) // 注册游戏服务端主服务控制相关路由
	}
	{
		// 需要鉴权的路由
		systemRouter.SysUserRouter.InitSysUserRouter(PrivateGroup) // 注册用户路由
	}

	return Router
}
