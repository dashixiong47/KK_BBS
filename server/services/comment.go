package services

import (
	"context"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"time"
)

type CommentData struct {
	Comments []Comment `json:"comments"`
	Total    int64     `json:"total"`
}
type Comment struct {
	ReplyComment
	Reply []ReplyComment `json:"reply"`
}
type ReplyComment struct {
	Id          string    `json:"id"`
	Content     string    `json:"content"`
	TopicId     string    `json:"topicId"`
	ParentId    string    `json:"parentId"`
	ReplyToUser string    `json:"replyToUser"`
	ReplyUser   any       `json:"replyUser"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Total       int64     `json:"total"`
	User        any       `json:"user"`
	Like        int       `json:"like"`
	LikeState   bool      `json:"likeState"`
}

// GetCommentLike 获取点赞数
func GetCommentLike(topicId, commentId, subCommentId int) int {
	ctx := context.Background()
	zName, name := GetCommentLikeName(topicId, commentId, subCommentId)

	like, err := db.Rdb.ZScore(ctx, zName, name).Result()
	if err != nil {
		like = 0
	}
	return int(like)
}

// GetCommentList 获取评论列表
func GetCommentList(topicId int, docs []Comment, comments, allReplyComments []models.Comment) []Comment {
	// 将子评论和它们的总数添加到相应的一级评论中
	for _, comment := range comments {
		var reply []ReplyComment
		var totalReplyComments int64
		db.DB.Model(&models.Comment{}).Where("parent_id = ?", comment.ID).Count(&totalReplyComments)

		for _, replyComment := range allReplyComments {
			if comment.ID == replyComment.ParentID {
				reply = append(reply, ReplyComment{
					Id:          db.GetStrID(replyComment.ID),
					Content:     replyComment.Content,
					TopicId:     db.GetStrID(replyComment.TopicID),
					ParentId:    db.GetStrID(replyComment.ParentID),
					ReplyToUser: db.GetStrID(replyComment.ReplyToUserID),
					User:        GetUserInfo(replyComment.UserID),
					ReplyUser:   GetUserInfo(replyComment.ReplyToUserID),
					CreatedAt:   replyComment.CreatedAt,
					UpdatedAt:   replyComment.UpdatedAt,
				})
			}
		}

		docs = append(docs, Comment{
			ReplyComment: ReplyComment{
				Id:          db.GetStrID(comment.ID),
				Content:     comment.Content,
				TopicId:     db.GetStrID(comment.TopicID),
				ParentId:    db.GetStrID(comment.ParentID),
				ReplyToUser: db.GetStrID(comment.ReplyToUserID),
				CreatedAt:   comment.CreatedAt,
				UpdatedAt:   comment.UpdatedAt,
				Total:       totalReplyComments,
				User:        GetUserInfo(comment.UserID),
			},
			Reply: reply,
		})
	}
	return docs
}

func UpDataLike(comments []Comment) []Comment {
	for i := range comments {
		comments[i].Like = GetCommentLike(db.GetIntID(comments[i].TopicId), db.GetIntID(comments[i].Id), 0)
		comments[i].LikeState = GetLikeState(db.GetIntID(comments[i].Id), 0)
		for j := range comments[i].Reply {
			comments[i].Reply[j].Like = GetCommentLike(db.GetIntID(comments[i].Reply[j].TopicId), db.GetIntID(comments[i].Id), db.GetIntID(comments[i].Reply[j].Id))
			comments[i].Reply[j].LikeState = GetLikeState(db.GetIntID(comments[i].Id), db.GetIntID(comments[i].Reply[j].Id))
		}
	}
	return comments
}

func GetCommentLikeName(topicId, commentId, subCommentId int) (string, string) {
	zName := ""
	name := ""
	if subCommentId == 0 {
		zName = fmt.Sprintf("comment_like:%d", topicId)
		name = fmt.Sprintf("%d", commentId)
	} else {
		zName = fmt.Sprintf("comment_like_sub:%d", topicId)
		name = fmt.Sprintf("%d", subCommentId)
	}
	return zName, name
}

func GetLikeState(commentID, subCommentID int) bool {

	var like int64
	db.DB.Model(models.CommentLike{}).
		Where("comment_id = ?", commentID).
		Where("sub_comment_id = ?", subCommentID).
		Count(&like)
	return like > 0
}

// CommentLickPlus 点赞
func CommentLickPlus(topicId, commentId, subCommentId int) error {
	ctx := context.Background()
	zName, name := GetCommentLikeName(topicId, commentId, subCommentId)
	_, err := db.Rdb.ZIncrBy(ctx, zName, 1, name).Result()
	if err != nil {
		return err
	}
	return nil
}

// CommentLickMinus 取消点赞
func CommentLickMinus(topicId, commentId, subCommentId int) error {
	ctx := context.Background()
	zName, name := GetCommentLikeName(topicId, commentId, subCommentId)
	_, err := db.Rdb.ZIncrBy(ctx, zName, -1, name).Result()
	if err != nil {
		return err
	}
	return nil
}

// IsCommentLike 是否点赞
func IsCommentLike(topicId, commentId, subCommentId int) bool {
	ctx := context.Background()
	zName, name := GetCommentLikeName(topicId, commentId, subCommentId)
	_, err := db.Rdb.ZRank(ctx, zName, name).Result()
	if err != nil {
		return false
	}
	if err == redis.Nil {
		return false
	}
	return true
}

// AddCommentIntegral 回复获得积分
func AddCommentIntegral(userId, topicId uint) error {
	return AddIntegral(userId, 1, 10, 2, topicId, "发表评论获得积分")
}

// ReplyCommentIntegral 被回复发帖人获得积分
func ReplyCommentIntegral(userId, topicId uint) error {

	return AddIntegral(userId, 1, 5, 3, topicId, "被回复获得积分")
}

// EarnPoints 评论获得积分
func EarnPoints(tx *gorm.DB, comment *models.Comment, userId uint) error {

	// 如果评论者是发帖人，则不获得积分
	if userId == comment.UserID {
		return nil
	}

	// 检查用户是否已经评论过该话题
	var exists bool
	err := tx.Model(&models.Comment{}).
		Select("1").
		Where("user_id = ? AND topic_id = ?", comment.UserID, comment.TopicID).
		Limit(1).
		Find(&exists).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	// 如果用户没有对该主题发表过评论，则进行积分奖励
	if !exists {
		// 回复获得积分
		if err = AddCommentIntegral(comment.UserID, comment.TopicID); err != nil {
			tx.Rollback()
			return err
		}
		// 发帖人获得积分
		if err = ReplyCommentIntegral(userId, comment.TopicID); err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}
