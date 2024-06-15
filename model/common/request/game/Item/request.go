package item

import (
	"gateway/model/common/request"
	"gateway/model/dnf/custom"
)

type ItemRequest struct {
	request.PageInfo
	custom.GameItem
}
