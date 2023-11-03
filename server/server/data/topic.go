package data

import (
	"context"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func GetTopicLikeName(topicId int) string {
	return fmt.Sprintf("topic_like:%d", topicId)
}

// CreateTopicLike 创建帖子点赞
func CreateTopicLike(topicId int) error {
	ctx := context.Background()
	_, err := db.Rdb.ZIncrBy(ctx, "topic_like", 0, fmt.Sprintf("%v", topicId)).Result()
	if err != nil {
		return err
	}
	return nil
}

// TopicLickPlus 帖子点赞数加一
func TopicLickPlus(topicId int) error {
	ctx := context.Background()
	_, err := db.Rdb.ZIncrBy(ctx, "topic_like", 1, fmt.Sprintf("%v", topicId)).Result()
	if err != nil {
		return err
	}
	return nil
}

// TopicLickMinus 帖子点赞数减一
func TopicLickMinus(topicId int) error {
	ctx := context.Background()
	_, err := db.Rdb.ZIncrBy(ctx, "topic_like", -1, fmt.Sprintf("%v", topicId)).Result()
	if err != nil {
		return err
	}
	return nil
}

// IsTopic 是否存在该帖子
func IsTopic(topicId int) bool {
	ctx := context.Background()
	_, err := db.Rdb.ZRank(ctx, "topic_like", fmt.Sprintf("%v", topicId)).Result()
	if err != nil || err == redis.Nil {
		return false
	}
	return true
}

// GetTopicLikeCount 获取帖子点赞数
func GetTopicLikeCount(topicId uint) int64 {
	ctx := context.Background()
	count, err := db.Rdb.ZScore(ctx, "topic_like", fmt.Sprintf("%v", topicId)).Result()
	if err != nil {
		return 0
	}
	return int64(count)
}

// GetTopicCommentCount 获取帖子评论数
func GetTopicCommentCount(topicId uint) int64 {
	ctx := context.Background()
	count, err := db.Rdb.ZCard(ctx, fmt.Sprintf("comment_like:%v", topicId)).Result()
	if err != nil {
		return 0
	}
	return count
}

// IsTopicLike 是否点赞
func IsTopicLike(topicId uint, userId uint) bool {
	var count int64
	db.DB.Model(&models.TopicLike{}).
		Where("topic_id = ?", topicId).
		Where("user_id = ?", userId).
		Count(&count)
	return count > 0
}

// TopicViewPlus 帖子浏览记录
func TopicViewPlus(topicId, userId uint) error {
	ctx := context.Background()
	_, err := db.Rdb.ZIncrBy(ctx, "topic_view", 1, fmt.Sprintf("%v", topicId)).Result()
	if err != nil {
		return err
	}
	view := models.TopicView{
		UserID:  userId,
		TopicID: topicId,
	}
	db.DB.Create(&view)
	return nil
}

// GetTopicViewCount 获取帖子浏览数
func GetTopicViewCount(topicId uint) int64 {
	ctx := context.Background()
	count, err := db.Rdb.ZScore(ctx, "topic_view", fmt.Sprintf("%v", topicId)).Result()
	if err != nil {
		return 0
	}
	return int64(count)
}

// GetTopicCount 获取帖子总数
func GetTopicCount() int64 {
	ctx := context.Background()
	count, err := db.Rdb.ZCard(ctx, "topic_like").Result()
	if err != nil {
		return 0
	}
	return count
}

const (
	cacheTTL = 2 * time.Hour // 定义缓存的过期时间为2小时
)

// GetTopicCollectCount 获取收藏帖子的数量
func GetTopicCollectCount(userId uint) int64 {
	name := fmt.Sprintf("topic_collect_count:%v", userId) // 格式化缓存键名
	ctx := context.Background()                           // 创建一个背景上下文

	// 尝试从Redis获取缓存数据
	result, err := db.Rdb.Get(ctx, name).Result()
	if err == nil {
		// 如果缓存数据存在，尝试解析为整数
		if count, err := strconv.ParseInt(result, 10, 64); err == nil {
			return count // 如果解析成功，返回数量
		}
	}

	// 如果缓存数据不存在或解析失败，从数据库获取数量
	var count int64
	err = db.DB.Model(&models.Collection{}).Where("user_id = ?", userId).Count(&count).Error
	if err == nil && count > 0 {
		// 将数量转换为字符串并存储到Redis，设置过期时间为2小时
		_ = db.Rdb.Set(ctx, name, count, cacheTTL).Err()
		return count // 返回从数据库获取的数量
	}

	return 0 // 如果有任何错误或没有数据，返回0
}

// ClearTopicCollectCountCache 清除收藏帖子的数量缓存
func ClearTopicCollectCountCache(userId uint) {
	name := fmt.Sprintf("topic_collect_count:%v", userId) // 格式化缓存键名
	ctx := context.Background()                           // 创建一个背景上下文
	_ = db.Rdb.Del(ctx, name).Err()                       // 删除缓存数据
}

// IsTopicCollect 是否收藏
func IsTopicCollect(topicId, userId uint) bool {
	var count int64
	db.DB.Model(&models.Collection{}).
		Where("topic_id = ?", topicId).
		Where("user_id = ?", userId).
		Count(&count)
	return count > 0
}

// IsTopicComment 是否评论
func IsTopicComment(topicId, userId uint) bool {
	var count int64
	db.DB.Model(&models.Comment{}).
		Where("topic_id = ?", topicId).
		Where("user_id = ?", userId).
		Count(&count)
	return count > 0
}

//// GetTopicCollect 获取收藏帖子
//func GetTopicCollect(userId uint) []models.Collection {
//	name := fmt.Sprintf("topic_collect:%v", userId) // 格式化缓存键名
//	ctx := context.Background()
//
//	// 尝试从Redis获取缓存数据
//	result, err := db.Rdb.Get(ctx, name).Result()
//	if err == nil {
//		var data []models.Collection
//		// 如果缓存数据存在，尝试解析为JSON
//		if json.Unmarshal([]byte(result), &data) == nil {
//			return data // 如果解析成功，返回数据
//		}
//	}
//
//	var docs []models.Collection
//	// 如果缓存数据不存在或解析失败，从数据库获取数据
//	if db.DB.Where("user_id = ?", userId).Find(&docs).Error == nil && len(docs) > 0 {
//		// 将数据序列化为JSON并存储到Redis
//		if marshal, err := json.Marshal(docs); err == nil {
//			// 将序列化后的数据存储到Redis，设置过期时间为2小时
//			_ = db.Rdb.Set(ctx, name, marshal, cacheTTL).Err()
//		}
//		return docs // 返回从数据库获取的数据
//	}
//
//	return []models.Collection{} // 如果有任何错误或没有数据，返回空数组
//}
