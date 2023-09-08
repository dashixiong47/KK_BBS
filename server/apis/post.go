package apis

import (
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/captcha"
	"github.com/gin-gonic/gin"
)

type Post struct {
	Ctx *gin.Context
}
type formData struct {
	models.Post
	CaptchaId string `json:"captchaId"`
	Code      string `json:"code"`
}

// PostCreate 创建帖子
func (p *Post) PostCreate() utils.ResponseData {
	var doc formData
	var postServer server.PostServer
	err := p.Ctx.ShouldBindJSON(&doc)
	if err != nil {

		return utils.JsonParameterError("json_error")
	}
	verifyCaptcha := captcha.VerifyCaptcha(doc.CaptchaId, doc.Code)
	if !verifyCaptcha {
		return utils.JsonFail("code_error")
	}
	var post = models.Post{
		UserID: doc.UserID,
		Title:  doc.Title,
		Tags:   doc.Tags,
		Covers: doc.Covers,
		Type:   doc.Type,
	}
	if doc.Type == 1 {
		var defaultPost = models.PostDefault{
			Content:       doc.DefaultPost.Content,
			Hidden:        doc.DefaultPost.Hidden,
			HiddenContent: doc.DefaultPost.HiddenContent,
		}
		post.DefaultPost = defaultPost
	}
	if doc.Type == 2 {
		var imagePost = models.PostImage{
			Introduction: doc.ImagePost.Introduction,
			ImageID:      doc.ImagePost.ImageID,
		}
		post.ImagePost = imagePost
	}
	if doc.Type == 3 {
		var videoPost = models.PostVideo{
			Introduction: doc.VideoPost.Introduction,
			VideoID:      doc.VideoPost.VideoID,
		}
		post.VideoPost = videoPost
	}
	if doc.Type == 4 {
		var textPost = models.PostText{
			Introduction: doc.TextPost.Introduction,
			TextID:       doc.TextPost.TextID,
		}
		post.TextPost = textPost
	}
	if err := postServer.Create(&post); err != nil {
		return utils.JsonFail(err.Error())
	}
	return utils.JsonSuccess(doc)
}

// GetList 帖子列表
func (p *Post) GetList() utils.ResponseData {
	var paging utils.Paging
	paging.GetPaging(p.Ctx)
	var postServer server.PostServer
	list, err := postServer.GetPostList(paging)
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(list)
}
