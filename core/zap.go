package core

import (
	"fmt"
	"gateway/core/internal"
	"gateway/global"
	"gateway/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
// Author [wangrui19970405](https://github.com/wangrui19970405)
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.GatewayConf.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.GatewayConf.Zap.Director)
		_ = os.Mkdir(global.GatewayConf.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.GatewayConf.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
