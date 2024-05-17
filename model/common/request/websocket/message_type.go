package websocket

var (
	// 心跳
	HeartBeat = 1000
	// 主要服务控制
	MainServiceControl = 1001
	// 主要服务控制返回
	MainServiceControlReturn = 1002
)

type MessageType struct {
	Type int         `json:"type"`
	Msg  interface{} `json:"msg"`
}
