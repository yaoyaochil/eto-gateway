package taiwan_billing_cera_request

import "gateway/model/common/request"

type CashCeraRequest struct {
	request.PageInfo
	Account string `json:"account"`
}
