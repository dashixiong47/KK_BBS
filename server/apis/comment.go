package apis

import (
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/middleware"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
)

type Comment struct {
	Ctx *gin.Context
}

// FormDataComment 评论
type FormDataComment struct {
	ParentID      string `json:"parentId"`                   // 父级ID
	TopicID       string `json:"topicId" binding:"required"` // 帖子ID
	ReplyToUserID string `json:"replyToUserId"`              // 被回复用户的ID
	Content       string `json:"content" binding:"required"` // 评论内容
}

// commentId
type commentId struct {
	CommentId    string `json:"commentId" binding:"required"`
	SubCommentId string `json:"subCommentId"`
}

// PostCreate 创建评论
func (c *Comment) PostCreate() utils.ResponseData {
	if authMiddleware, data := middleware.AuthMiddleware(c.Ctx); !authMiddleware {
		return *data
	}
	var formData FormDataComment
	var comment models.Comment
	err := c.Ctx.ShouldBindJSON(&formData)
	if err != nil {
		return utils.JsonParameterError("json_error")
	}
	//
	if formData.Content == "" {
		return utils.JsonParameterError("content_error")
	}
	comment.Content = formData.Content
	// 帖子ID
	topicId := db.GetIntID(formData.TopicID)
	if topicId == 0 {
		return utils.JsonParameterError("topic_id_error")
	}
	comment.TopicID = uint(topicId)
	// 父级ID
	if formData.ParentID != "" {
		id := db.GetIntID(formData.ParentID)
		if id == 0 {
			return utils.JsonParameterError("parent_id_error")
		}
		comment.ParentID = uint(id)
	}
	// 被回复用户的ID
	if formData.ReplyToUserID != "" {
		id := db.GetIntID(formData.ReplyToUserID)
		if id == 0 {
			return utils.JsonParameterError("reply_to_user_id_error")
		}
		comment.ReplyToUserID = uint(id)
	}
	// user_id
	userId := utils.UserIDUint(c.Ctx)
	comment.UserID = userId
	var commentServer server.CommentServer
	doc, err := commentServer.Create(&comment)
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(doc)
}

// GetListBy 获取评论列表
func (c *Comment) GetListBy(id string) utils.ResponseData {
	//if authMiddleware, services := middleware.AuthMiddleware(c.Ctx); !authMiddleware {
	//	return *services
	//}
	_type := c.Ctx.DefaultQuery("type", "all")
	var paging utils.Paging
	paging.GetPaging(c.Ctx)
	var commentServer server.CommentServer
	intID := db.GetIntID(id)
	doc, err := commentServer.GetList(intID, paging, _type)
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(doc)
}

// GetSubListBy 获取subComment列表
func (c *Comment) GetSubListBy(id string) utils.ResponseData {
	//if authMiddleware, services := middleware.AuthMiddleware(c.Ctx); !authMiddleware {
	//	return *services
	//}
	topicId := c.Ctx.DefaultQuery("topicId", "")
	if topicId == "" {
		return utils.JsonParameterError("topic_id_error")
	}
	var paging utils.Paging
	paging.GetPaging(c.Ctx)
	var commentServer server.CommentServer
	intSubID := db.GetIntID(id)
	intTopicId := db.GetIntID(topicId)

	doc, err := commentServer.GetSubCommentList(intTopicId, intSubID, paging)
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(doc)
}

// PostLikeBy 点赞
func (c *Comment) PostLikeBy(id string) utils.ResponseData {
	if authMiddleware, data := middleware.AuthMiddleware(c.Ctx); !authMiddleware {
		return *data
	}
	var err error
	userId := utils.UserIDInt(c.Ctx)
	var commentId commentId
	err = c.Ctx.ShouldBindJSON(&commentId)
	if err != nil {
		return utils.JsonParameterError("json_error")
	}
	var commentServer server.CommentServer
	intID := db.GetIntID(id)
	err = commentServer.LikeComment(intID, userId, db.GetIntID(commentId.CommentId), db.GetIntID(commentId.SubCommentId))
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(nil)
}
