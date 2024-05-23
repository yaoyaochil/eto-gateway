package game

import (
	"gateway/global"
	"gateway/model/dnf/custom"
)

type ItemService struct {
}

// CreateItem 创建Item
func (s *ItemService) CreateItem(data custom.GameItem) (err error) {
	err = global.GatewayDBs["system"].Create(&data).Error
	return err
}
