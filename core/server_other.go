//go:build !windows
// +build !windows

package core

import (
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

// initServer 初始化服务器
func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
