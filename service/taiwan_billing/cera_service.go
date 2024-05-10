package taiwan_billing

import (
	"errors"
	"fmt"
	"gateway/global"
	"gateway/model/common/request/taiwan_billing_cera_request"
	d_taiwan "gateway/model/dnf/d_taiwan"
	model "gateway/model/dnf/taiwan_billing"
	"time"
)

type CeraService struct {
}

// GetCera 获取账号点券
func (c *CeraService) GetCera(info model.CashCera) (data model.CashCera, err error) {
	err = global.GatewayDBs["taiwan_billing"].Where("account = ?", info.Account).First(&data).Error
	return data, err
}

// RechargeCera 充值点券
func (c *CeraService) RechargeCera(info model.CashCera) (data model.CashCera, err error) {
	var cera model.CashCera
	if err = global.GatewayDBs["taiwan_billing"].Where("account = ?", info.Account).First(&cera).Error; err != nil {
		return data, err
	}
	cera.Cera += info.Cera
	// 获取当前时间
	currentTime := time.Now()
	// 格式化成 ISO 8601 格式
	formattedTime := currentTime.Format("2006-01-02T15:04:05-07:00")
	// 将格式化后的时间字符串解析为时间类型
	parsedTime, err := time.Parse("2006-01-02T15:04:05-07:00", formattedTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	cera.ModDate = parsedTime
	err = global.GatewayDBs["taiwan_billing"].Save(&cera).Error
	return cera, err
}

type BatchRechargeCeraRequest struct {
	Ids  []string `json:"ids"`
	Cera int32    `json:"cera"`
}

// BatchRechargeCera 批量充值点券
func (c *CeraService) BatchRechargeCera(info BatchRechargeCeraRequest) (err error) {
	err = global.GatewayDBs["taiwan_billing"].Where("account IN ?", info.Ids).Update("cera", info.Cera).Error
	return err
}

// DeductCera 扣减点券
func (c *CeraService) DeductCera(info model.CashCera) (data model.CashCera, err error) {
	var cera model.CashCera
	if err = global.GatewayDBs["taiwan_billing"].Where("account = ?", info.Account).First(&cera).Error; err != nil {
		return data, err
	}
	// 如果点券不足
	if cera.Cera < info.Cera {
		return data, errors.New("点券不足")
	}

	cera.Cera -= info.Cera
	// 获取当前时间
	currentTime := time.Now()
	// 格式化成 ISO 8601 格式
	formattedTime := currentTime.Format("2006-01-02T15:04:05-07:00")
	// 将格式化后的时间字符串解析为时间类型
	parsedTime, err := time.Parse("2006-01-02T15:04:05-07:00", formattedTime)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	cera.ModDate = parsedTime
	err = global.GatewayDBs["taiwan_billing"].Save(&cera).Error
	return cera, err
}

// 分页获取账号点券
func (c *CeraService) GetCeraList(info taiwan_billing_cera_request.CashCeraRequest) (list []model.CashCera, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GatewayDBs["taiwan_billing"]

	if info.Account != "" {
		var account d_taiwan.Account
		err = global.GatewayDBs["d_taiwan"].Where("accountname = ?", info.Account).First(&account).Error
		if err != nil {
			return list, total, err
		}
		db = db.Where("account = ?", fmt.Sprintf("%d", account.UID))
	}

	err = db.Limit(limit).Offset(offset).Find(&list).Error
	db.Model(&model.CashCera{}).Count(&total)
	return list, total, err
}
