package internal

import (
	"fmt"
	"gateway/global"

	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
	dbName string
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer, dbName string) *writer {
	return &writer{Writer: w, dbName: dbName}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch w.dbName {
	case "GormMysqlDTaiwan":
		logZap = global.GatewayConf.GormMysqlDTaiwan.LogZap
	case "GormMysqlSystem":
		logZap = global.GatewayConf.GormMysqlSystem.LogZap
	default:
		logZap = false
	}
	if logZap {
		global.GatewayLog.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
