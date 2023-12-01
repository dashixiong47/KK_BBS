package router

import (
	"encoding/json"
	"github.com/dashixiong47/KK_BBS/utils/jwt"
	"github.com/gin-gonic/gin"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

// UserConn 用户连接信息结构体
type UserConn struct {
	user  int
	conn  *websocket.Conn
	_chan chan Message
}

// Message 消息结构体
type Message struct {
	Type int             `json:"type"`
	Data json.RawMessage `json:"data"`
}

// MsgData 消息数据结构体
type MsgData struct {
	Type      int    `json:"type"`
	UserId    int    `json:"userId"`
	TimeStamp int64  `json:"timeStamp"`
	Data      string `json:"data"`
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
		log.Printf("WebSocket升级失败: %v\n", err)
		return
	}
	defer conn.Close()

	ch := make(chan Message, 100)
	defer close(ch)

	go handleMessages(conn, ch)

	userConns.Store(userId, UserConn{
		user:  userId,
		conn:  conn,
		_chan: ch,
	})

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("读取消息失败: %v\n", err)
			break
		}
		var msgData MsgData
		err = json.Unmarshal(message, &msgData)
		if err != nil {
			log.Printf("JSON解析失败: %v\n", err)
			continue
		}
		if msgData.TimeStamp != 0 {
			ping := map[string]int64{
				"timeStamp": time.Now().UnixNano(),
			}
			marshal, _ := json.Marshal(ping)
			SendMsg(userId, Message{Type: messageType, Data: marshal})
			return
		}
		if msgData.UserId != userId {
			SendMsg(msgData.UserId, Message{Type: messageType, Data: message})
		}
	}
}

// handleMessages 处理并发送消息
func handleMessages(conn *websocket.Conn, ch chan Message) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("handleMessages中出现panic: %v\n", err)
		}
	}()

	for msg := range ch {
		if err := conn.WriteMessage(websocket.TextMessage, msg.Data); err != nil {
			log.Printf("消息发送失败: %v\n", err)
			break
		}
	}
}

// SendMsg 发送消息到指定的用户
func SendMsg(userId int, msg Message) {
	if val, ok := userConns.Load(userId); ok {
		userConn := val.(UserConn)
		userConn._chan <- msg
	}
}
