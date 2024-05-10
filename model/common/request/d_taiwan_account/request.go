package d_taiwan_account

import (
	"gateway/model/common/request"
	model "gateway/model/dnf/d_taiwan"
)

type AccountRequest struct {
	request.PageInfo
	model.Account
}
