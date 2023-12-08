package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/dashixiong47/KK_BBS/data"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"github.com/go-redis/redis/v8"
	"time"
)

type MessageServer struct {
}

func (m *MessageServer) MessageInfo(ctx context.Context, userId int, paging utils.Paging) (interface{}, error) {
	var docs []models.Messages
	var count int64

	// 以 user_id 进行筛选
	tx := db.DB.Model(&models.Messages{}).Where("user_id = ?", userId)

	// 首先按照 is_view 排序（未读消息优先），然后按照 id 降序排序
	tx = tx.Order("is_view asc, id desc")

	// 计算总消息数量
	if err := tx.Count(&count).Error; err != nil {
		return nil, err
	}

	// 应用分页
	tx = paging.SetPaging(tx)

	// 执行查询
	if err := tx.Find(&docs).Error; err != nil {
		return nil, err
	}

	var docMap = make([]map[string]any, 0, len(docs))
	for _, doc := range docs {
		content, err := getContent(ctx, doc)
		if err != nil {
			klog.Error("Error getting message content: %v", err)
			continue // 或者考虑返回错误
		}
		docMap = append(docMap, content)
	}

	return map[string]any{
		"list":  docMap,
		"total": count,
	}, nil
}

// UnreadMessageInfo 获取未读消息数量
func (m *MessageServer) UnreadMessageInfo(userId int) (interface{}, error) {
	var count int64
	tx := db.DB.Model(&models.Messages{})
	tx.Where("user_id = ? and is_view = ?", userId, 0)
	tx.Count(&count)
	return count, nil
}

// ReadBy 单个已读
func (m *MessageServer) ReadBy(messageId, userId uint) error {
	// 更新 Messages 表中指定用户的指定消息的 is_view 字段
	err := db.DB.Model(&models.Messages{}).
		Where("user_id = ?", userId). // 筛选特定用户
		Where("id = ?", messageId).   // 筛选特定消息
		Update("is_view", 1).Error    // 将 is_view 更新为 1，表示已读

	// 检查并返回任何可能发生的错误
	if err != nil {
		return errors.New("update_message_error")
	}

	// 没有错误时返回 nil
	return nil
}

// ReadAll 全部已读
func (m *MessageServer) ReadAll(userId int) error {
	return db.DB.Model(&models.Messages{}).Where("user_id = ?", userId).Update("is_view", 1).Error

}

func getContent(ctx context.Context, docs models.Messages) (map[string]any, error) {
	var content = map[string]any{
		"id":        db.GetStrID(docs.ID),
		"type":      docs.Type,
		"isView":    docs.IsView,
		"createdAt": docs.CreatedAt,
	}
	if docs.RelatedUserID != 0 {
		content["relatedUser"] = data.GetUserInfo(docs.RelatedUserID)
	}
	switch docs.Type {
	case 1:
		// 系统通知
		content["title"] = "系统通知"
		content["content"] = "您有一条新的系统通知"
		return content, nil
	case 2, 3, 4, 6:
		topicTitle := getTopicName(ctx, docs.SupplementID)
		content["content"] = docs.SupplementData
		content["detail"] = fmt.Sprintf("/topic/%v", db.GetStrID(docs.SupplementID))
		content["topicTitle"] = topicTitle
		return content, nil
	default:
		// 未知的消息类型
		return nil, fmt.Errorf("unknown message type: %v", docs.Type)
	}
}

func getTopicName(ctx context.Context, topicId int) string {
	// 定义缓存键
	cacheKey := fmt.Sprintf("topic:name:%v", topicId)

	// 尝试从 Redis 获取
	nickname, err := db.Rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		return nickname
	}
	if !errors.Is(err, redis.Nil) {
		// 记录非缓存未命中的错误
		klog.Error("Error retrieving from Redis: %v", err)
	}

	// 从数据库获取
	var topic models.Topic
	err = db.DB.WithContext(ctx).Model(&topic).
		Where("id = ?", topicId).
		Select("title").
		Take(&topic).Error
	if err != nil {
		return "帖子找不到啦" // 返回数据库查询错误
	}

	// 将结果设置回 Redis
	_, err = db.Rdb.Set(ctx, cacheKey, topic.Title, 2*time.Hour).Result()
	if err != nil {
		// 记录 Redis 设置操作的错误
		klog.Error("Error setting Redis key: %v", err)
	}

	return topic.Title
}
