package taiwan_cain

import (
	"gateway/global"
	"gateway/model/common/request/taiwan_cain_request"
	model "gateway/model/dnf/taiwan_cain"
)

type TaiwanCainCharacterInfoService struct {
}

// GetCharacterInfo 获取角色信息
func (t *TaiwanCainCharacterInfoService) GetCharacterInfo(info model.CharacInfo) (data model.CharacInfo, err error) {
	err = global.GatewayDBs["taiwan_cain"].Where("charac_no = ? AND delete_flag = ?", info.CharacNo, 0).Find(&data).Error
	return data, err
}

// GetCharacterInfoList 获取角色信息列表
func (t *TaiwanCainCharacterInfoService) GetCharacterInfoList(info taiwan_cain_request.TaiwanCainCharacterInfoRequest) (list []model.CharacInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GatewayDBs["taiwan_cain"]

	if info.MID != 0 {
		db = db.Where("m_id = ?", info.MID)
	}

	err = db.Limit(limit).Offset(offset).Find(&list).Error
	db.Model(&model.CharacInfo{}).Count(&total)
	return list, total, err
}
