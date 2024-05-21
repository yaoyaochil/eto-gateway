package gamecontrol

import (
	"fmt"
	"gateway/model/common/request"
	"gateway/model/common/response"
	"gateway/service"
	"gateway/source/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChannelListAPI struct{}

var channelService = service.GroupServiceApp.GameControl.ChannelService

// GetChannelList 获取频道列表
func (cl *ChannelListAPI) GetChannelList(c *gin.Context) {
	var request request.PageInfo
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dir := "/home/neople/game/cfg/"
	pre := "siroco"
	data, total, err := channelService.GetPagedConfigs(dir, pre, request.PageSize, request.Page)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(response.PageResult{
		List:     data,
		Total:    total,
		Page:     request.Page,
		PageSize: request.PageSize,
	}, c)
}

// UpdateChannel 更新频道
func (cl *ChannelListAPI) UpdateChannel(c *gin.Context) {
	var request template.ChannelConfig
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dir := "/home/neople/game/cfg/"
	err := channelService.UpdateConfig(dir, request.GCNo, request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// CreateChannel 创建频道
func (cl *ChannelListAPI) CreateChannel(c *gin.Context) {
	var request template.ChannelConfig
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dir := "/home/neople/game/cfg/"
	err := channelService.CreateConfig(request, dir)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteChannel 删除频道
func (cl *ChannelListAPI) DeleteChannel(c *gin.Context) {
	var request template.ChannelConfig
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	dir := "/home/neople/game/cfg/"
	err := channelService.DeleteConfig(dir, request.GCNo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// StartChannelServiceByGCNo 启动频道
func (cl *ChannelListAPI) StartChannelServiceByGCNo(c *gin.Context) {
	var request template.ChannelConfig
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := channelService.StartChannelServiceByGCNo(request.GCNo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("启动成功", c)
}

// StopChannelServiceByGCNo 停止频道
func (cl *ChannelListAPI) StopChannelServiceByGCNo(c *gin.Context) {
	var request template.ChannelConfig
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := channelService.StopChannelServiceByGCNo(request.GCNo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("停止成功", c)
}

// RestartChannelServiceByGCNo 重启频道
func (cl *ChannelListAPI) RestartChannelServiceByGCNo(c *gin.Context) {
	var request template.ChannelConfig
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := channelService.RestartChannelServiceByGCNo(request.GCNo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("重启成功", c)
}

// GetChannelStatusByGCNo 获取频道状态
func (cl *ChannelListAPI) GetChannelStatusByGCNo(c *gin.Context) {
	var request template.ChannelConfig
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	status, err := channelService.GetChannelStatusByGCNo(request.GCNo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(status, c)
}

type ChannelStatusAPI struct {
	Ids []int `json:"ids"`
}

// GetChannelStatusByGCNos 批量获取频道状态
func (cl *ChannelListAPI) GetChannelStatusByGCNos(c *gin.Context) {
	var request ChannelStatusAPI
	if err := c.ShouldBindJSON(&request); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	status, err := channelService.GetChannelStatusByGCNos(request.Ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(status, c)
}

var upgrader = websocket.Upgrader{
	// 允许跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	// 读取缓冲区大小
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// 实时获取频道运行日志
func (cl *ChannelListAPI) GetChannelLogByGCNo(c *gin.Context) {
	id := c.Query("id")
	// 如果id为空，返回错误
	if id == "" {
		response.FailWithMessage("id不能为空", c)
		return
	}
	path := fmt.Sprintf("/home/neople/game/log/siroco%s/*.init", id)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	defer conn.Close()

	channelService.GetChannelRunLog(path, conn)
}
