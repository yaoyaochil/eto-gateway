package gamecontrol

import "gateway/global"

type ServerInfo struct {
	global.GateWayMODEL
	ServerId     int    `json:"server_id"`     // 服务器id
	ServerName   string `json:"server_name"`   // 服务器名称
	ServerIp     string `json:"server_ip"`     // 服务器ip
	ServerPort   int    `json:"server_port"`   // 服务器端口
	ServerStatus int    `json:"server_status"` // 服务器状态
	ServerGroup  int    `json:"server_group"`  // 服务器组
}
