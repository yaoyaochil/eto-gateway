package d_taiwan

import (
	"gateway/model/common/request/d_taiwan_account"
	"gateway/model/common/response"
	model "gateway/model/dnf/d_taiwan"
	"gateway/protobuf/service/protobuf/game"
	"gateway/service"
	"gateway/service/d_taiwan"
	"gateway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
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
 * @Description: DnfLoginProtobuf protobuf协议登录
 * @receiver a AccountApi
 * @param c
 */
func (a *AccountApi) DnfLoginProtobuf(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		response.FailWithMessage("获取失败", c)
		c.Abort()
		return
	}

	// 调用protobuf服务 获取登录信息
	var loginRequest game.LoginRequest
	if err := proto.Unmarshal(data, &loginRequest); err != nil {
		response.FailWithMessage("登录失败", c)
		c.Abort()
		return
	}

	account, err := accountService.GetAccount(model.Account{Accountname: loginRequest.Accountname})
	if err != nil {
		response.FailWithMessage("账号不存在或密码错误", c)
		c.Abort()
		return
	}
	if !utils.VerifyDNFPassword(loginRequest.Password, account.Password) {
		response.FailWithMessage("账号不存在或密码错误", c)
		c.Abort()
		return
	}

	base64Content, err := utils.EncryptBase64WithUID(account.UID)
	if err != nil {
		response.FailWithMessage("登录失败,请联系客服", c)
		c.Abort()
		return
	}

	loginRes := &game.LoginResponse{
		Code: 0,
		Msg:  "登录成功",
		Data: base64Content,
	}
	data, err = proto.Marshal(loginRes)
	if err != nil {
		response.FailWithMessage("登录失败", c)
		c.Abort()
		return
	}
	c.Data(http.StatusOK, "application/x-protobuf", data)
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

type CreateAccountRequest struct {
	CashCera      int32 `json:"cashCera"`      // 0
	CashCeraPoint int32 `json:"cashCeraPoint"` // 0
	model.Account
}

/**
 * @Description: CreateAccount 创建账号
 * @receiver a AccountApi
 * @param c
 */
func (a *AccountApi) CreateAccount(c *gin.Context) {
	var body CreateAccountRequest
	_ = c.ShouldBindJSON(&body)
	_, err := accountService.GetAccount(body.Account)
	if err == nil {
		response.FailWithMessage("账号已存在", c)
		c.Abort()
		return
	}
	_, err = accountService.CreateAccount(d_taiwan.CreateAccountRequest(body))
	if err != nil {
		response.FailWithMessage("创建失败", c)
		c.Abort()
		return
	}
	response.OkWithMessage("创建成功", c)
}

/**
 * @Description: Register 注册
 * @receiver a AccountApi
 * @param c
 */
func (a *AccountApi) Register(c *gin.Context) {
	var body CreateAccountRequest
	_ = c.ShouldBindJSON(&body)
	_, err := accountService.GetAccount(body.Account)
	if err == nil {
		response.FailWithMessage("账号已存在", c)
		c.Abort()
		return
	}
	body.Password = utils.CalculateMD5Hash(body.Password)
	_, err = accountService.CreateAccount(d_taiwan.CreateAccountRequest(body))
	if err != nil {
		response.FailWithMessage("注册失败", c)
		c.Abort()
		return
	}
	response.OkWithMessage("注册成功", c)
}

/**
 * @Description: GetAccountList 获取账号列表
 * @receiver a AccountApi
 * @param c
 */
func (a *AccountApi) GetAccountList(c *gin.Context) {
	var body d_taiwan_account.AccountRequest
	_ = c.ShouldBindJSON(&body)
	data, total, err := accountService.GetAccountList(body)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		c.Abort()
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     data,
		Total:    total,
		Page:     body.Page,
		PageSize: body.PageSize,
	}, "获取成功", c)
}

// ResetLimitCreateCharacter 重置创建角色限制
func (a *AccountApi) ResetLimitCreateCharacter(c *gin.Context) {
	var body model.Account
	_ = c.ShouldBindJSON(&body)
	err := accountService.ResetLimitCreateCharacter(body.UID)
	if err != nil {
		response.FailWithMessage("重置失败", c)
		c.Abort()
		return
	}
	response.OkWithMessage("重置成功", c)
}
