package service

import (
	"gateway/service/d_taiwan"
	"gateway/service/game"
	"gateway/service/game_control"
	"gateway/service/system"
	"gateway/service/taiwan_billing"
	"gateway/service/taiwan_cain"
	"gateway/service/taiwan_cain_2nd"
)

type GroupService struct {
	DTaiwan            d_taiwan.GroupDTaiwanService
	TaiwanBilling      taiwan_billing.GroupDTaiwanBillingService
	TaiwanCain         taiwan_cain.GroupTaiwanCainService
	TaiwanCain2ND      taiwan_cain_2nd.GroupService
	SystemServiceGroup system.ServiceGroup
	GameControl        game_control.GroupService
	Game               game.GroupService
}

var GroupServiceApp = new(GroupService)
