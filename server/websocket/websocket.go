package websocket

import (
	"encoding/json"
	"github.com/dashixiong47/KK_BBS/utils/jwt"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"github.com/gin-gonic/gin"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

// UserConn 用户连接信息结构体
type UserConn struct {
	UserId     int
	Conn       *websocket.Conn
	Chan       chan Message
	LastActive time.Time
}

// Message 消息结构体
type Message struct {
	Type int             `json:"type"`
	Data json.RawMessage `json:"data"`
}

// MsgData 消息数据结构体
type MsgData struct {
	Type   int    `json:"type"`
	UserId int    `json:"userId"`
	Ping   int64  `json:"ping"`
	Data   string `json:"data"`
}

// ConnManager 管理WebSocket连接
type ConnManager struct {
	userConns sync.Map
}

// 全局变量
var GlobalConnManager = NewConnManager()

// NewConnManager 创建新的连接管理器实例
func NewConnManager() *ConnManager {
	return &ConnManager{}
}

// AddConn 添加新连接
func (m *ConnManager) AddConn(userId int, conn *websocket.Conn) {
	// 检查是否已存在该用户的连接
	if val, ok := m.userConns.LoadAndDelete(userId); ok {
		// 关闭旧连接
		oldUserConn := val.(UserConn)
		oldUserConn.Conn.Close()
		close(oldUserConn.Chan)
	}

	// 添加新连接
	ch := make(chan Message, 100)
	m.userConns.Store(userId, UserConn{
		UserId:     userId,
		Conn:       conn,
		Chan:       ch,
		LastActive: time.Now(),
	})

	// 启动消息处理协程
	go handleMessages(conn, ch)
}

// RemoveConn 移除连接
func (m *ConnManager) RemoveConn(userId int) {
	if val, ok := m.userConns.Load(userId); ok {
		userConn := val.(UserConn)
		userConn.Conn.Close()
		close(userConn.Chan)
		m.userConns.Delete(userId)
	}
}

// SendMsg 发送消息给特定用户
func (m *ConnManager) SendMsg(userId int, msg Message) {
	if val, ok := m.userConns.Load(userId); ok {
		userConn := val.(UserConn)
		userConn.Chan <- msg
	}
}

// UpdateLastActive 更新用户的最后活动时间
func (m *ConnManager) UpdateLastActive(userId int) {
	if val, ok := m.userConns.Load(userId); ok {
		userConn := val.(UserConn)
		userConn.LastActive = time.Now()
		m.userConns.Store(userId, userConn)
	}
}

// CheckTimeout 检查和断开超时的连接
func (m *ConnManager) CheckTimeout(duration time.Duration) {
	now := time.Now()
	m.userConns.Range(func(key, value interface{}) bool {
		userConn := value.(UserConn)
		if now.Sub(userConn.LastActive) > duration {
			userConn.Conn.Close()
			m.RemoveConn(userConn.UserId)
		}
		return true
	})
}

// CountOnlineUsers 返回当前在线用户数量
func (m *ConnManager) CountOnlineUsers() int {
	var count int
	m.userConns.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}

// upgrade 用于WebSocket升级
var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有源
	},
}

// WebsocketHandler 处理WebSocket请求
func WebsocketHandler(c *gin.Context) {
	header := c.Query("Authorization")
	if header == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	user, err := jwt.ParseToken(header)
	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	userId := int(user.Claims.(jwtv5.MapClaims)["id"].(float64))

	conn, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		klog.Error("WebSocket升级失败: %v\n", err)
		return
	}
	defer conn.Close()

	// 使用全局ConnManager实例添加新连接
	GlobalConnManager.AddConn(userId, conn)

	// 心跳检测定时器
	heartbeatTimer := time.NewTicker(10 * time.Second)
	defer func() {
		heartbeatTimer.Stop()
		GlobalConnManager.RemoveConn(userId)
	}()

	for {
		select {
		case <-heartbeatTimer.C:
			GlobalConnManager.CheckTimeout(20 * time.Second)
		default:
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				klog.Error("websocketHandler", err)
				break
			}
			var msgData MsgData
			err = json.Unmarshal(message, &msgData)
			if err != nil {
				klog.Error("解析JSON失败:%v", err)
				continue
			}

			if msgData.Ping != 0 {
				ping := map[string]int64{
					"pong": time.Now().Unix(),
				}
				marshal, _ := json.Marshal(ping)
				GlobalConnManager.SendMsg(userId, Message{Type: messageType, Data: marshal})
				GlobalConnManager.UpdateLastActive(userId)
			}

		}
	}
}

// handleMessages 处理并发送消息
func handleMessages(conn *websocket.Conn, ch chan Message) {
	defer func() {
		if err := recover(); err != nil {
			klog.Error("handleMessages", err)
		}
	}()

	for msg := range ch {
		if err := conn.WriteMessage(websocket.TextMessage, msg.Data); err != nil {
			klog.Error("handleMessages", err)
			break
		}
	}
}
