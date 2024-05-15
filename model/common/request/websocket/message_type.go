package websocket

var (
	// 心跳
	HeartBeat = 1000
	// 服务控制
	ServiceControl = 1001
)

type MessageType struct {
	Type int         `json:"type"`
	Msg  interface{} `json:"msg"`
}
