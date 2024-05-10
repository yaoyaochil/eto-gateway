package taiwan_cain_2nd

import (
	v1 "gateway/api/v1"
	"github.com/gin-gonic/gin"
)

type EmailRouter struct{}

func (e *EmailRouter) InitEmailRouter(router *gin.RouterGroup) {
	router = router.Group("email")
	var emailApi = v1.ApiGroupApp.TaiwanCain2NDApiGroup.EmailApi
	{
		router.POST("sendEmail", emailApi.SendEmail)
	}
}
