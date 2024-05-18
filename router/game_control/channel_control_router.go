package gamecontrol

import (
	v1 "gateway/api/v1"

	"github.com/gin-gonic/gin"
)

type ChannelRouter struct{}

func (c *ChannelRouter) InitChannelRouter(Router *gin.RouterGroup) {
	ChannelRouter := Router.Group("channel")
	var channelApi = v1.ApiGroupApp.GameControlApiGroup.ChannelListAPI
	{
		ChannelRouter.POST("getChannelList", channelApi.GetChannelList)
		ChannelRouter.POST("updateChannel", channelApi.UpdateChannel)
		ChannelRouter.POST("createChannel", channelApi.CreateChannel)
		ChannelRouter.POST("deleteChannel", channelApi.DeleteChannel)
	}
	{
		ChannelRouter.POST("startChannel", channelApi.StartChannelServiceByGCNo)
		ChannelRouter.POST("stopChannel", channelApi.StopChannelServiceByGCNo)
		ChannelRouter.POST("restartChannel", channelApi.RestartChannelServiceByGCNo)
		ChannelRouter.POST("getChannelStatus", channelApi.GetChannelStatusByGCNo)
		ChannelRouter.POST("GetChannelStatusByGCNos", channelApi.GetChannelStatusByGCNos)
		ChannelRouter.GET("getChannelLog", channelApi.GetChannelLogByGCNo)
	}
}
