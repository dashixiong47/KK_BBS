package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/server/data"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"time"
)

type CommentServer struct {
}

// Create 创建评论
func (s *CommentServer) Create(comment *models.Comment) (*models.Comment, error) {
	tx := db.DB.Begin()
	// 评论积分
	err := data.EarnPoints(tx, comment)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 创建评论
	if err := tx.Create(comment).Error; err != nil {
		tx.Rollback()
		return nil, errors.New("create_comment_error")
	}
	// 更新评论点赞排行
	ctx := context.Background()
	var zName, name string
	if comment.ParentID == 0 {
		zName, name = data.GetCommentLikeName(int(comment.TopicID), int(comment.ID), 0)
	} else {
		zName, name = data.GetCommentLikeName(int(comment.TopicID), int(comment.ParentID), int(comment.ID))
	}

	if _, err = db.Rdb.ZIncrBy(ctx, zName, 0, name).Result(); err != nil {
		tx.Rollback()
		return nil, errors.New("create_comment_error")
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		return nil, errors.New("create_comment_error")
	}

	// 删除对应的 Redis 缓存
	if err = db.Del(fmt.Sprintf("comment_list:%v", comment.TopicID)); err != nil {
		klog.Error("删除 Redis 缓存失败: %v", err)
	}

	return comment, nil
}

// GetList 获取评论列表
func (s *CommentServer) GetList(topicId int, paging utils.Paging, _type string) (any, error) {
	name := fmt.Sprintf("comment_list:%v:p:%v:s:%v:%v", topicId, paging.Page, paging.PageSize, _type)
	ctx := context.Background()
	var response data.CommentData // 最终返回的数据
	// 尝试从 Redis 中获取缓存
	result, err := db.Rdb.Get(ctx, name).Result()
	if err == nil {
		_ = json.Unmarshal([]byte(result), &response)
		response.Comments = data.UpDataLike(response.Comments)
		return response, nil
	}

	var comments = make([]models.Comment, 0)         // 一级评论
	var allReplyComments = make([]models.Comment, 0) // 所有子评论
	var docs = make([]data.Comment, 0)               // 评论列表

	// 获取一级评论总数
	var totalPrimaryComments int64
	db.DB.Model(&models.Comment{}).Where("topic_id = ?", topicId).Where("parent_id = 0").Count(&totalPrimaryComments)

	if _type == "hot" {
		// 查询一级热评
		// 计算起始和结束的排名（0-based index）
		start := int64((paging.Page - 1) * paging.PageSize)
		end := start + int64(paging.PageSize) - 1

		// 查询有序集合（按分数从大到小排序）
		topIds, err := db.Rdb.ZRevRange(ctx, fmt.Sprintf("comment_like:%d", topicId), start, end).Result()
		if err != nil {
			fmt.Println("Error:", err)
		}
		err = db.DB.Where("id in ?", topIds).Find(&comments).Error
		if err != nil {
			return nil, errors.New("get_comments_error")
		}
	} else {
		// 查询一级评论
		err = paging.SetPaging(db.DB).Where("topic_id = ?", topicId).
			Where("parent_id = 0").
			Order("created_at desc").
			Find(&comments).Error
		if err != nil {
			return nil, errors.New("get_comments_error")
		}
	}

	// 收集所有子评论
	for _, comment := range comments {
		var replyCommentsForOneParent []models.Comment
		err = db.DB.Where("parent_id = ?", comment.ID).
			Limit(3).
			Order("created_at desc").
			Find(&replyCommentsForOneParent).Error
		if err != nil {
			return nil, errors.New("get_comments_error")
		}
		allReplyComments = append(allReplyComments, replyCommentsForOneParent...)
	}

	// 缓存结果到 Redis 并返回
	response = data.CommentData{
		Comments: data.GetCommentList(topicId, docs, comments, allReplyComments),
		Total:    totalPrimaryComments,
	}
	// 缓存到 Redis
	marshal, err := json.Marshal(response)
	if err != nil {
		klog.Error("comments_redis_error: %v", err)
	}
	if err := db.Rdb.Set(ctx, name, marshal, time.Hour).Err(); err != nil {
		klog.Error("comments_redis_error: %v", err)
	}
	response.Comments = data.UpDataLike(response.Comments)
	return response, nil
}

// GetSubCommentList 获取子评论列表
func (s *CommentServer) GetSubCommentList(topicId, subId int, paging utils.Paging) (any, error) {
	var comment []models.Comment
	err := paging.SetPaging(db.DB).Model(comment).
		Where("topic_id = ?", topicId).
		Where("parent_id = ?", subId).
		Order("created_at desc").
		Find(&comment).Error
	if err != nil {
		return nil, errors.New("get_comments_error")
	}
	var docs []map[string]interface{}
	for _, v := range comment {
		docs = append(docs, map[string]interface{}{
			"id":          db.GetID(v.ID),
			"content":     v.Content,
			"topicId":     db.GetID(v.TopicID),
			"parentId":    db.GetID(v.ParentID),
			"replyToUser": db.GetID(v.ReplyToUserID),
			"user":        data.GetUserInfo(v.UserID),
			"replyUser":   data.GetUserInfo(v.ReplyToUserID),
			"createdAt":   v.CreatedAt,
			"updatedAt":   v.UpdatedAt,
		})

	}
	return docs, nil
}

// LikeComment 点赞
func (s *CommentServer) LikeComment(topicId, userId, commentId, subCommentId int) error {
	var commentLike models.CommentLike
	//state := data.IsCommentLike(topicId, commentId, subCommentId)
	//if !state {
	//	return errors.New("is_not_comment")
	//}
	// 查找是否存在该点赞记录（包括软删除的记录）
	tex := db.DB.Unscoped().
		Where("user_id = ? AND comment_id = ?", userId, commentId)
	if subCommentId != 0 {
		tex.Where("sub_comment_id = ?", subCommentId)
	}
	tex.First(&commentLike)
	// 开启数据库事务
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 如果存在该记录
	if commentLike.ID != 0 && commentLike.DeletedAt.Valid {
		// 如果是软删除的记录，则恢复
		if err := tx.Model(&commentLike).
			Unscoped().
			Update("deleted_at", nil).Error; err != nil {
			tx.Rollback()
			return errors.New("comment_like_error")
		}
		err := data.CommentLickPlus(topicId, commentId, subCommentId)
		if err != nil {
			tx.Rollback()
			return errors.New("comment_like_error")
		}
		//	如果是正常的记录，则进行软删除
	} else if commentLike.ID != 0 && !commentLike.DeletedAt.Valid {
		if err := tx.Delete(&commentLike).Error; err != nil {
			tx.Rollback()
			return errors.New("comment_like_error")
		}
		err := data.CommentLickMinus(topicId, commentId, subCommentId)
		if err != nil {
			tx.Rollback()
			return errors.New("comment_like_error")
		}
		// 不存在该记录，创建新记录
	} else {
		// 不存在该记录，创建新记录
		commentLike = models.CommentLike{
			UserID:       uint(userId),
			CommentID:    uint(commentId),
			SubCommentID: uint(subCommentId),
		}
		if err := tx.Create(&commentLike).Error; err != nil {
			tx.Rollback()
			return errors.New("comment_like_error")
		}
		err := data.CommentLickPlus(topicId, commentId, subCommentId)
		if err != nil {
			tx.Rollback()
			return errors.New("comment_like_error")
		}
	}

	// 提交数据库事务
	if err := tx.Commit().Error; err != nil {
		return errors.New("comment_like_error")
	}

	return nil
}
