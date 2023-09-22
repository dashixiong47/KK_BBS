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
	value, _ := c.Ctx.Get("id")
	comment.UserID = uint(value.(float64))
	var commentServer server.CommentServer
	doc, err := commentServer.Create(&comment)
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(doc)
}

// GetListBy 获取评论列表
func (c *Comment) GetListBy(id string) utils.ResponseData {
	//if authMiddleware, data := middleware.AuthMiddleware(c.Ctx); !authMiddleware {
	//	return *data
	//}
	var paging utils.Paging
	paging.GetPaging(c.Ctx)
	var commentServer server.CommentServer
	intID := db.GetIntID(id)
	doc, err := commentServer.GetList(intID, paging)
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(doc)
}