package core

import (
	"fmt"
	"gateway/global"
	"gateway/initialize"
	"gateway/service/system"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 初始化redis服务
	//initialize.RedisInit()

	// 从db加载jwt数据
	if global.GatewayDBs["system"] != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%s", global.GatewayConf.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)

	fmt.Printf(`ETO网关服务启动成功%s 
`, address)
	global.GatewayLog.Error(s.ListenAndServe().Error())
}
