package initialize

import (
	"gateway/global"
	"gateway/utils"

	"github.com/songzhibin97/gkit/cache/local_cache"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.GatewayConf.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.GatewayConf.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.GatewayCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}
