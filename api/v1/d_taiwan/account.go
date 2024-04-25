package d_taiwan

import (
	"gateway/model/common/response"
	model "gateway/model/dnf/d_taiwan"
	"gateway/service"
	"gateway/utils"
	"github.com/gin-gonic/gin"
)

type AccountApi struct{}

var accountService = service.GroupServiceApp.DTaiwan.AccountService

/**
 * @Description: GetAccount 根据账号名获取账号信息
 * @receiver a AccountApi
 * @param c
 */
func (a *AccountApi) GetAccount(c *gin.Context) {
	var body model.Account
	_ = c.ShouldBindJSON(&body)
	data, err := accountService.GetAccount(body)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		c.Abort()
		return
	}
	response.OkWithData(data, c)
}

/**
 * @Description: DnfLogin 登录
 * @receiver a AccountApi
 * @param c
 */
func (a *AccountApi) DnfLogin(c *gin.Context) {
	var body model.Account
	_ = c.ShouldBindJSON(&body)
	data, err := accountService.GetAccount(body)
	if err != nil {
		response.FailWithMessage("账号不存在或密码错误", c)
		c.Abort()
		return
	}
	if !utils.VerifyDNFPassword(body.Password, data.Password) {
		response.FailWithMessage("账号不存在或密码错误", c)
		c.Abort()
		return
	}
	base64Content, err := utils.EncryptBase64WithUID(data.UID)
	if err != nil {
		response.FailWithMessage("登录失败,请联系客服", c)
		c.Abort()
		return
	}
	response.OkWithLogin(base64Content, c)
}
