package message

import (
	"encoding/json"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"github.com/dashixiong47/KK_BBS/websocket"
)

var msgData = map[int]string{
	1: "您有一条系统消息",
	2: "您有一条评论消息",
	3: "您有一条回复消息",
	4: "您有一条点赞消息",
	5: "您有一条@消息",
	6: "您有一条收藏消息",
	7: "您有一条关注消息",
}

func SendMessage(message models.Messages) {
	err := db.DB.Create(&message).Error
	if err != nil {
		klog.Error(err.Error())
	}
	var msg = map[string]any{
		"type":     message.Type,
		"services": msgData[message.Type],
	}
	marshal, err := json.Marshal(msg)
	if err != nil {
		klog.Error(err.Error())
	}
	websocket.GlobalConnManager.SendMsg(int(message.UserID), websocket.Message{
		Type: 1,
		Data: marshal,
	})
}

// SendCommentMessage 发送评论消息
func SendCommentMessage(relatedUserID, masterUserId, topicId uint, content string) error {
	var message models.Messages
	message.UserID = masterUserId
	message.Type = 2
	message.RelatedUserID = relatedUserID
	message.SupplementID = int(topicId)
	message.SupplementData = content
	SendMessage(message)
	return nil
}

// SendReplyMessage 发送回复消息
func SendReplyMessage(relatedUserID, masterUserId, topicId uint, content string) error {
	var message models.Messages
	message.UserID = masterUserId
	message.Type = 3
	message.RelatedUserID = relatedUserID
	message.SupplementID = int(topicId)
	message.SupplementData = content
	SendMessage(message)
	return nil
}

// SendLikeMessage  发送点赞消息
func SendLikeMessage(relatedUserID, masterUserId, topicId uint, content string) error {
	var message models.Messages
	message.UserID = masterUserId
	message.Type = 4
	message.RelatedUserID = relatedUserID
	message.SupplementID = int(topicId)
	message.SupplementData = content
	SendMessage(message)
	return nil
}

// SendCollectMessage 发送回复消息
func SendCollectMessage(relatedUserID, masterUserId, topicId uint, content string) error {
	var message models.Messages
	message.UserID = masterUserId
	message.Type = 6
	message.RelatedUserID = relatedUserID
	message.SupplementID = int(topicId)
	message.SupplementData = content
	SendMessage(message)
	return nil
}
