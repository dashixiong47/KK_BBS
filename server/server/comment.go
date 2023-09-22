package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"time"
)

type CommentServer struct {
}

// Create 创建评论
func (s *CommentServer) Create(comment *models.Comment) (*models.Comment, error) {
	err := db.DB.Create(comment).Error
	if err != nil {
		return nil, errors.New("create_comment_error")
	}
	if err := db.Del(fmt.Sprintf("comment_list:%v", comment.TopicID)); err != nil {
		klog.Error("删除 Redis 缓存失败: %v", err)
	}
	return comment, nil
}

func (s *CommentServer) GetList(topicId int, paging utils.Paging) (any, error) {
	ctx := context.Background()

	// 尝试从 Redis 中获取缓存
	result, err := db.Rdb.Get(ctx, fmt.Sprintf("comment_list:%v:p:%v:s:%v", topicId, paging.Page, paging.PageSize)).Result()
	if err == nil {
		var data map[string]interface{}
		_ = json.Unmarshal([]byte(result), &data)
		return data, nil
	}

	var comments []models.Comment
	var docs []map[string]interface{}
	// 查询一级评论
	err = paging.SetPaging(db.DB).Where("topic_id = ?", topicId).
		Where("parent_id = 0").
		Order("created_at desc").
		Find(&comments).Error
	if err != nil {
		return nil, errors.New("get_comments_error")
	}

	// 获取一级评论总数
	var totalPrimaryComments int64
	db.DB.Model(&models.Comment{}).Where("topic_id = ?", topicId).Where("parent_id = 0").Count(&totalPrimaryComments)

	var ids []uint
	for _, comment := range comments {
		ids = append(ids, comment.ID)
	}

	// 收集所有子评论
	var allReplyComments []models.Comment
	for _, id := range ids {
		var replyCommentsForOneParent []models.Comment
		err = db.DB.Where("parent_id = ?", id).
			Limit(3).
			Order("created_at desc").
			Find(&replyCommentsForOneParent).Error
		if err != nil {
			return nil, errors.New("get_comments_error")
		}
		allReplyComments = append(allReplyComments, replyCommentsForOneParent...)
	}

	// 将子评论和它们的总数添加到相应的一级评论中
	for _, comment := range comments {
		var reply []map[string]interface{}
		var totalReplyComments int64
		db.DB.Model(&models.Comment{}).Where("parent_id = ?", comment.ID).Count(&totalReplyComments)

		for _, replyComment := range allReplyComments {
			if comment.ID == replyComment.ParentID {
				reply = append(reply, map[string]interface{}{
					"id":          db.GetID(replyComment.ID),
					"content":     replyComment.Content,
					"topicId":     db.GetID(replyComment.TopicID),
					"parentId":    db.GetID(replyComment.ParentID),
					"replyToUser": db.GetID(replyComment.ReplyToUserID),
					"user":        getUserInfo(replyComment.UserID),
					"replyUser":   getUserInfo(replyComment.ReplyToUserID),
					"createdAt":   replyComment.CreatedAt,
					"updatedAt":   replyComment.UpdatedAt,
				})
			}
		}

		docs = append(docs, map[string]interface{}{
			"id":          db.GetID(comment.ID),
			"content":     comment.Content,
			"topicId":     db.GetID(comment.TopicID),
			"parentId":    db.GetID(comment.ParentID),
			"replyToUser": db.GetID(comment.ReplyToUserID),
			"createdAt":   comment.CreatedAt,
			"updatedAt":   comment.UpdatedAt,
			"total":       totalReplyComments,
			"user":        getUserInfo(comment.UserID),
			"reply":       reply,
		})
	}

	// 缓存结果到 Redis 并返回
	response := map[string]interface{}{
		"comments": docs,
		"total":    totalPrimaryComments,
	}

	marshal, err := json.Marshal(response)
	if err != nil {
		return nil, errors.New("json_marshal_error")
	}
	if err := db.Rdb.Set(ctx, fmt.Sprintf("comment_list:%v:p:%v:s:%v", topicId, paging.Page, paging.PageSize), marshal, time.Hour).Err(); err != nil {
		return nil, errors.New("redis_set_error")
	}

	return response, nil
}

func getUserInfo(userId uint) any {
	ctx := context.Background()
	userData, err := db.Rdb.Get(ctx, fmt.Sprintf("user_info:%v", userId)).Result()
	if err == nil {
		var data map[string]interface{}
		_ = json.Unmarshal([]byte(userData), &data)
		return data
	} else {

		var user models.User
		db.DB.Where("id = ?", userId).Find(&user)
		var data = map[string]interface{}{
			"id":       db.GetID(user.ID),
			"avatar":   user.Avatar,
			"username": user.Username,
			"nickname": user.Nickname,
		}
		db.Rdb.Set(ctx, fmt.Sprintf("user_info:%v", user.ID), data, time.Hour*2)

		return data
	}
}
