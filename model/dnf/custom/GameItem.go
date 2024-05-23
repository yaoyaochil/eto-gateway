package custom

import "gateway/global"

type GameItem struct {
	global.GateWayMODEL
	ItemName  string `json:"itemName" gorm:"comment:物品名称;type:varchar(255);not null;default:''"` // 物品名称
	ItemType  int    `json:"itemType" gorm:"comment:物品类型;type:int;not null;default:0"`           // 0:装备 1:道具 2:材料
	ItemLevel int    `json:"itemLevel" gorm:"comment:物品等级;type:int;not null;default:0"`          // 0:白 1:绿 2:蓝 3:紫 4:史诗 5:传说 6:神话
	ItemPrice int    `json:"itemPrice" gorm:"comment:物品价格;type:int;not null;default:0"`          // 价格
	ItemDesc  string `json:"itemDesc" gorm:"comment:物品描述;type:varchar(255);not null;default:''"` // 物品描述
}
