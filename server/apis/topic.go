package apis

import (
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/middleware"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/captcha"
	"github.com/dashixiong47/KK_BBS/utils/html"
	"github.com/gin-gonic/gin"
)

type Topic struct {
	Ctx *gin.Context
}

type formData struct {
	models.Topic
	//必填
	CaptchaId string `json:"captchaId" binding:"required"`
	Code      string `json:"code" binding:"required"`
}

// PostCreate 创建帖子
func (t *Topic) PostCreate() utils.ResponseData {
	if authMiddleware, data := middleware.AuthMiddleware(t.Ctx); !authMiddleware {
		return *data
	}
	var doc formData
	var postServer server.TopicServer
	err := t.Ctx.ShouldBindJSON(&doc)
	if err != nil {

		return utils.JsonParameterError("json_error")
	}
	verifyCaptcha := captcha.VerifyCaptcha(doc.CaptchaId, doc.Code)
	if !verifyCaptcha {
		return utils.JsonFail("code_error")
	}
	var post = models.Topic{
		UserID: doc.UserID,
		Title:  doc.Title,
		Tags:   doc.Tags,
		Covers: doc.Covers,
		Type:   doc.Type,
	}
	if doc.Type == 1 {
		summary, err := html.GetHtmlSummary(doc.TopicBasic.Content)
		if err != nil {
			return utils.JsonFail(err)
		}
		post.Summary = summary
		if doc.TopicBasic.HiddenContent != "" {
			doc.TopicBasic.Hidden = 1
		}
		var defaultPost = models.TopicBasic{
			Content:       doc.TopicBasic.Content,
			Hidden:        doc.TopicBasic.Hidden,
			HiddenContent: doc.TopicBasic.HiddenContent,
		}
		post.TopicBasic = defaultPost
	}
	if doc.Type == 2 {
		var imagePost = models.TopicImage{
			Introduction: doc.TopicImage.Introduction,
			ImageID:      doc.TopicImage.ImageID,
		}
		post.TopicImage = imagePost
	}
	if doc.Type == 3 {
		var videoPost = models.TopicVideo{
			Introduction: doc.TopicVideo.Introduction,
			VideoID:      doc.TopicVideo.VideoID,
		}
		post.TopicVideo = videoPost
	}
	if doc.Type == 4 {
		var textPost = models.TopicText{
			Introduction: doc.TopicText.Introduction,
			TextID:       doc.TopicText.TextID,
		}
		post.TopicText = textPost
	}
	id, err := postServer.Create(&post)
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(id)
}

// GetList 帖子列表
func (t *Topic) GetList() utils.ResponseData {
	//if authMiddleware, data := middleware.AuthMiddleware(t.Ctx); !authMiddleware {
	//	return *data
	//}
	var paging utils.Paging
	paging.GetPaging(t.Ctx)
	var topicServer server.TopicServer
	list, err := topicServer.GetPostList(paging)
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(list)
}

// GetBy 帖子详情
func (t *Topic) GetBy(id string) utils.ResponseData {
	//if authMiddleware, data := middleware.AuthMiddleware(t.Ctx); !authMiddleware {
	//	return *data
	//}
	var topicServer server.TopicServer
	detail, err := topicServer.GetTopicDetail(db.GetIntID(id))
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(detail)
}
