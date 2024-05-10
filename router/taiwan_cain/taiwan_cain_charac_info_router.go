package taiwan_cain

import (
	v1 "gateway/api/v1"
	"github.com/gin-gonic/gin"
)

type CharacInfoRouter struct {
}

func (c *CharacInfoRouter) InitCharacInfoRouter(router *gin.RouterGroup) {
	router = router.Group("charac_info")
	var api = v1.ApiGroupApp.TaiwanCainApiGroup.CharacInfoApi
	{
		router.POST("getCharacInfo", api.GetCharacInfo)
		router.POST("getCharacInfoList", api.GetCharacInfoList)
	}
}
