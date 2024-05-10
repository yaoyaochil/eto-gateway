package taiwan_cain_2nd_request

import model "gateway/model/dnf/taiwan_cain_2nd"

type SendEmailRequest struct {
	CharNo       int32  `json:"charNo"`       // 角色编号
	SendCharName string `json:"sendCharName"` // 发送者角色名
	LetterText   string `json:"letterText"`   // 信件内容
	model.Postal
}
