package taiwan_cain_2nd

import (
	"gateway/global"
	"gateway/model/common/request/taiwan_cain_2nd_request"
	model "gateway/model/dnf/taiwan_cain_2nd"
	"gateway/utils"
)

type EmailService struct {
}

// SendEmail 发送邮件
func (e *EmailService) SendEmail(info taiwan_cain_2nd_request.SendEmailRequest) (err error) {
	now := utils.GetDofTime()
	// TODO: 操作letter表
	var letter model.Letter
	letter = model.Letter{
		CharacNo:       info.CharNo,
		SendCharacNo:   0,
		SendCharacName: info.SendCharName,
		RegDate:        now, // 发送时间
		LetterText:     info.LetterText,
		Stat:           1, // 1未读 2已读 3存储
	}

	err = global.GatewayDBs["taiwan_cain_2nd"].Create(&letter).Error
	if err != nil {
		return err
	}

	// TODO: 操作postal表
	var postal model.Postal
	postal = model.Postal{
		OccTime:         now,                  // 发送时间
		SendCharacNo:    0,                    // 发送者角色编号 0表示系统
		SendCharacName:  info.SendCharName,    // 发送者角色名
		ReceiveCharacNo: info.CharNo,          // 收件人角色编号
		ItemID:          info.ItemID,          // 道具ID
		AddInfo:         info.AddInfo,         // 道具数量
		LetterID:        letter.LetterID,      // 信件ID
		SeperateUpgrade: info.SeperateUpgrade, // 武器锻造等级
		Upgrade:         info.Upgrade,         // 武器强化等级
		AmplifyOption:   info.AmplifyOption,   // 武器附魔等级
		AmplifyValue:    info.AmplifyValue,    // 武器附魔值
		Gold:            info.Gold,            // 金币
		SealFlag:        info.SealFlag,        // 封印标志 0未封印 1封印
		RandomOption:    []byte{},             // 随机属性
		ItemGUID:        []byte{},             // 道具GUID
		//ReceiveTime:     now,                   // 领取时间
	}

	err = global.GatewayDBs["taiwan_cain_2nd"].Create(&postal).Error

	return err
}
