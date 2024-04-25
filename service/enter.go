package service

import (
	"gateway/service/d_taiwan"
	"gateway/service/system"
)

type GroupService struct {
	DTaiwan            d_taiwan.GroupDTaiwanService
	SystemServiceGroup system.ServiceGroup
}

var GroupServiceApp = new(GroupService)
