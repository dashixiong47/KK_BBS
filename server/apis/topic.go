package apis

import (
	"errors"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/middleware"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/captcha"
	"github.com/dashixiong47/KK_BBS/utils/html"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"github.com/gin-gonic/gin"
	"log"
)

type Topic struct {
	Ctx *gin.Context
}

type attachment struct {
	Coins       int    `json:"coins"`       // 金币
	Hidden      int    `json:"hidden"`      // 是否隐藏 评论可见
	Type        int    `json:"type"`        // 附件类型 1:文件 2:网盘
	Order       int    `json:"order"`       // 排序
	Name        string `json:"name"`        // 文件名
	FileType    string `json:"fileType"`    // 文件类型
	FileSize    int    `json:"fileSize"`    // 文件大小
	FileUrl     string `json:"fileUrl"`     // 地址
	NetDiskType string `json:"netDisk"`     // 网盘类型
	NetDiskUrl  string `json:"netDiskUrl"`  // 网盘地址
	NetDiskCode string `json:"netDiskCode"` // 网盘提取码
	Remake      string `json:"remake"`      // 备注
}

type formData struct {
	models.Topic
	Attachment []attachment `json:"attachment"`
	//必填
	CaptchaId string `json:"captchaId" binding:"required"`
	Code      string `json:"code" binding:"required"`
}

// PostCreate 用于创建帖子
func (t *Topic) PostCreate() utils.ResponseData {
	// 验证用户是否已通过身份验证
	if authMiddleware, data := middleware.AuthMiddleware(t.Ctx); !authMiddleware {
		return *data
	}

	// 从上下文中获取用户ID
	id, exists := t.Ctx.Get("id")
	if !exists {
		klog.Error("ID not found in context")
		return utils.JsonFail("id_error")
	}

	// 将ID从interface{}类型安全转换为float64，然后转换为uint
	idFloat, ok := id.(float64)
	if !ok {
		klog.Error("Failed to convert ID to float64")
		return utils.JsonFail("id_error")
	}
	uintId := uint(idFloat)

	// 解析请求体中的JSON数据到formData结构体
	var doc formData
	err := t.Ctx.ShouldBindJSON(&doc)
	if err != nil {
		klog.Error("PostCreate ShouldBindJSON", err)
		return utils.JsonParameterError("json_error")
	}
	doc.UserID = uintId
	verifyCaptcha := captcha.VerifyCaptcha(doc.CaptchaId, doc.Code)
	if !verifyCaptcha {
		return utils.JsonFail("code_error")
	}
	// 验证帖子类型是否合法
	if doc.Type < 1 || doc.Type > 4 {
		return utils.JsonParameterError("invalid_type")
	}

	// 初始化一个models.Topic结构体，然后调用createPostByType函数来根据不同类型的帖子填充结构体字段
	var post models.Topic
	err = t.createPostByType(&doc, &post)
	if err != nil {
		klog.Error("Failed to create post by type", err)
		return utils.JsonFail("type_error")
	}

	// 处理和验证附件数据
	attachments, err := t.handleAttachments(&doc, uintId)
	if err != nil {
		klog.Error("Failed to handle attachments", err)
		return utils.JsonFail("attachment_error")
	}
	// 创建帖子
	var topicServer server.TopicServer
	topicId, err := topicServer.Create(&post, &attachments)
	if err != nil {
		klog.Error("Failed to create topic", err)
		return utils.JsonFail(err)
	}

	// 返回成功的响应，包括新创建的帖子的ID
	return utils.JsonSuccess(topicId)
}

// createPostByType 根据不同的帖子类型创建相应的帖子模型
func (t *Topic) createPostByType(doc *formData, post *models.Topic) error {
	// 公共字段赋值
	post.UserID = doc.UserID
	post.Title = doc.Title
	post.Tags = doc.Tags
	post.Covers = doc.Covers
	post.Type = doc.Type
	post.GroupID = doc.GroupID

	// 根据帖子类型，填充特定字段
	switch doc.Type {
	case 1: // 基本帖子类型
		summary, err := html.GetHtmlSummary(doc.TopicBasic.Content)
		if err != nil {
			return err
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
	case 2: // 图片帖子类型
		var imagePost = models.TopicImage{
			Introduction: doc.TopicImage.Introduction,
			Images:       doc.TopicImage.Images,
		}
		post.TopicImage = imagePost
	case 3: // 视频帖子类型
		var videoPost = models.TopicVideo{
			Introduction: doc.TopicVideo.Introduction,
			VideoID:      doc.TopicVideo.VideoID,
		}
		post.TopicVideo = videoPost
	case 4: // 文本帖子类型
		log.Println(doc.TopicText.Texts)
		var textPost = models.TopicText{
			Introduction: doc.TopicText.Introduction,
			Texts:        doc.TopicText.Texts,
		}
		post.TopicText = textPost
	default:
		return errors.New("unsupported_post_type")
	}
	return nil
}

// handleAttachments 处理帖子的附件
func (t *Topic) handleAttachments(doc *formData, uintId uint) ([]models.Attachment, error) {
	var attachments []models.Attachment
	// 遍历附件数据
	for _, v := range doc.Attachment {
		var attachment = models.Attachment{
			UserID: uintId,
			Coins:  v.Coins,
			Hidden: v.Hidden,
			Type:   v.Type,
			Order:  v.Order,
			Name:   v.Name,
			Remake: v.Remake,
		}
		// 根据附件类型，处理特定字段
		if v.Type == 1 { // 文件附件
			if v.FileUrl == "" || v.FileType == "" || v.FileSize == 0 {
				continue // 无效的文件附件数据，跳过
			}
			attachment.FileSize = v.FileSize
			attachment.FileType = v.FileType
			attachment.FileUrl = v.FileUrl
		}
		if v.Type == 2 { // 网盘附件
			if v.NetDiskUrl == "" || v.NetDiskType == "" {
				continue // 无效的网盘附件数据，跳过
			}
			attachment.NetDiskType = v.NetDiskType
			attachment.NetDiskUrl = v.NetDiskUrl
			attachment.NetDiskCode = v.NetDiskCode
		}
		// 将处理好的附件数据添加到附件列表中
		attachments = append(attachments, attachment)
	}

	return attachments, nil
}

// GetList 帖子列表
func (t *Topic) GetList() utils.ResponseData {
	middleware.AuthMiddleware(t.Ctx)
	var _type = t.Ctx.DefaultQuery("type", "")
	var userIdStr = t.Ctx.DefaultQuery("userId", "0")
	var groupIDStr = t.Ctx.DefaultQuery("groupId", "0")

	userId := uint(db.GetIntID(userIdStr))
	var paging utils.Paging
	paging.GetPaging(t.Ctx)
	var selfUserId uint
	selfUserIdFlot, _ := t.Ctx.Get("id")
	if selfUserIdFlot == nil {
		selfUserId = uint(float64(0))
	} else {
		selfUserId = uint(selfUserIdFlot.(float64))
	}
	var topicServer server.TopicServer
	list, err := topicServer.GetTopicList(_type, groupIDStr, userId, selfUserId, paging)
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(list)
}

// GetBy 帖子详情
func (t *Topic) GetBy(topicId string) utils.ResponseData {
	middleware.AuthMiddleware(t.Ctx)
	var topicServer server.TopicServer
	userId, _ := t.Ctx.Get("id")
	if userId == nil {
		userId = float64(0)
	}
	detail, err := topicServer.GetTopicDetail(db.GetIntID(topicId), int(userId.(float64)))
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(detail)
}

// PostLikeBy 点赞帖子
func (t *Topic) PostLikeBy(topicId string) utils.ResponseData {
	if authMiddleware, data := middleware.AuthMiddleware(t.Ctx); !authMiddleware {
		return *data
	}
	userId, _ := t.Ctx.Get("id")
	var topicServer server.TopicServer

	err := topicServer.LikeTopic(db.GetIntID(topicId), int(userId.(float64)))
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(nil)
}

// PostCollectBy 收藏帖子
func (t *Topic) PostCollectBy(topicId string) utils.ResponseData {
	if authMiddleware, data := middleware.AuthMiddleware(t.Ctx); !authMiddleware {
		return *data
	}
	userId, _ := t.Ctx.Get("id")
	var topicServer server.TopicServer
	err := topicServer.CollectTopic(db.GetIntID(topicId), int(userId.(float64)))
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(nil)
}
