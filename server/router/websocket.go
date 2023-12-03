package router

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
	UserId int
	Conn   *websocket.Conn
	Chan   chan Message
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

// userConns 存储所有用户的连接信息
var userConns sync.Map

// upgrade 用于WebSocket升级
var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有源
	},
}

// websocketHandler 处理WebSocket请求
func websocketHandler(c *gin.Context) {
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

	ch := make(chan Message, 100)

	go handleMessages(conn, ch)

	userConns.Store(userId, UserConn{
		UserId: userId,
		Conn:   conn,
		Chan:   ch,
	})

	for {
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
				"ping": time.Now().Unix(),
			}
			marshal, _ := json.Marshal(ping)
			SendMsg(userId, Message{Type: messageType, Data: marshal})

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

// SendMsg 发送消息到指定的用户
func SendMsg(userId int, msg Message) {
	if val, ok := userConns.Load(userId); ok {
		userConn := val.(UserConn)
		userConn.Chan <- msg
	}
}
