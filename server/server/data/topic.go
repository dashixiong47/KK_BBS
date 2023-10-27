package data

import (
	"context"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/go-redis/redis/v8"
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
