package d_taiwan

import (
	"gateway/global"
	model "gateway/model/dnf/d_taiwan"
)

type AccountService struct {
}

// GetAccount 根据账号名获取账号信息
func (a *AccountService) GetAccount(body model.Account) (data model.Account, err error) {
	err = global.GatewayDBs["d_taiwan"].Where("accountname = ?", body.Accountname).First(&data).Error
	return data, err
}
