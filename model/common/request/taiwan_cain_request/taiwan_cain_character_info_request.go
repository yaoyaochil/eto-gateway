package taiwan_cain_request

import (
	"gateway/model/common/request"
	model "gateway/model/dnf/taiwan_cain"
)

type TaiwanCainCharacterInfoRequest struct {
	request.PageInfo
	model.CharacInfo
}
