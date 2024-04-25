package global

import (
	"gateway/config"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	GatewayLog                *zap.Logger
	GatewayCache              local_cache.Cache
	GatewayConf               config.Server
	GatewayVp                 *viper.Viper
	GatewayConcurrencyControl = &singleflight.Group{}
	GatewayDBs                map[string]*gorm.DB
)
