package game

import (
	"gateway/global"
	item "gateway/model/common/request/game/Item"
	"gateway/model/dnf/custom"
)

type ItemService struct {
}

// CreateItem 创建Item
func (s *ItemService) CreateItem(data custom.GameItem) (err error) {
	err = global.GatewayDBs["system"].Create(&data).Error
	return err
}

// GetItem 根据ID获取Item
func (s *ItemService) GetItem(id int) (data custom.GameItem, err error) {
	err = global.GatewayDBs["system"].Where("id = ?", id).First(&data).Error
	return data, err
}

// DeleteItem 根据ID删除Item
func (s *ItemService) DeleteItem(id int) (err error) {
	err = global.GatewayDBs["system"].Where("id = ?", id).Delete(&custom.GameItem{}).Error
	return err
}

// UpdateItem 更新Item
func (s *ItemService) UpdateItem(data custom.GameItem) (err error) {
	err = global.GatewayDBs["system"].Save(&data).Error
	return err
}

// GetItemList 分页获取Item列表
func (s *ItemService) GetItemList(info item.ItemRequest) (list []custom.GameItem, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GatewayDBs["system"]
	var itemList []custom.GameItem
	err = db.Model(&custom.GameItem{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	err = db.Limit(limit).Offset(offset).Find(&itemList).Error
	return itemList, total, err
}

// 解析Excel文件 并根据文件内容批量插入数据
func (s *ItemService) BatchCreateItem(data []custom.GameItem) (err error) {
	err = global.GatewayDBs["system"].Create(&data).Error
	return err
}
