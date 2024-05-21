package websocket

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ServiceControlApi struct {
	Clients       map[*websocket.Conn]bool
	LastHeartbeat map[*websocket.Conn]time.Time // 最后一次心跳时间
	mu            sync.Mutex                    // 添加互斥锁
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func (s *ServiceControlApi) UpgradeServiceControlApi(c *gin.Context) {
	// 初始化 Clients map 如果不初始化会Panics
	if s.Clients == nil {
		s.Clients = make(map[*websocket.Conn]bool)
	}

	if s.LastHeartbeat == nil {
		s.LastHeartbeat = make(map[*websocket.Conn]time.Time)
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	s.mu.Lock() // 加锁
	s.Clients[conn] = true
	s.LastHeartbeat[conn] = time.Now()
	s.mu.Unlock() // 解锁

	// 心跳检测
	go func() {
		for {
			time.Sleep(5 * time.Second)
			s.mu.Lock()
			lastHeartbeat, ok := s.LastHeartbeat[conn]
			s.mu.Unlock()

			// 如果没有心跳或者超过30秒没有心跳则关闭连接
			if !ok || time.Since(lastHeartbeat) > 30*time.Second {
				conn.Close()
				break
			}
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		// 更新心跳时间
		s.Heartbeat(conn)
		s.BroadcastMessage(msg, conn)
	}
}

// BroadcastMessage 广播消息 除了发送者以外的所有客户端
func (s *ServiceControlApi) BroadcastMessage(message []byte, sender *websocket.Conn) {
	s.mu.Lock()         // 加锁
	defer s.mu.Unlock() // 在函数退出时解锁

	for client := range s.Clients {
		if client != sender {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				client.Close()
				delete(s.Clients, client)
			}
		}
	}
}

// Heartbeat 心跳
func (s *ServiceControlApi) Heartbeat(conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.LastHeartbeat[conn] = time.Now()
}
