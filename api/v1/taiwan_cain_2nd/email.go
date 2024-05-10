package taiwan_cain_2nd

import (
	"gateway/model/common/request/taiwan_cain_2nd_request"
	"gateway/model/common/response"
	"gateway/service"
	"github.com/gin-gonic/gin"
)

type EmailApi struct {
}

var email_service = service.GroupServiceApp.TaiwanCain2ND.EmailService

// SendEmail 发送邮件
func (e *EmailApi) SendEmail(c *gin.Context) {
	var info taiwan_cain_2nd_request.SendEmailRequest
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = email_service.SendEmail(info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("发送成功", c)
}
