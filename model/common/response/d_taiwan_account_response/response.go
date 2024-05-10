package d_taiwan_account_response

import model "gateway/model/dnf/d_taiwan"

type AccountResponse struct {
	model.Account
	CashCera      int32 `json:"cashCera"`      // 0
	CashCeraPoint int32 `json:"cashCeraPoint"` // 0
	IsHack        int32 `json:"isHack"`        // 是否被封号 0否 1是
}
