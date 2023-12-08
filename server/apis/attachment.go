package apis

import (
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/middleware"
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
)

type Attachment struct {
	Ctx *gin.Context
}

// GetListBy GetList 获取附件列表
func (c *Attachment) GetListBy(topId string) utils.ResponseData {
	middleware.AuthMiddleware(c.Ctx)
	topicIdInt := db.GetIntID(topId)

	var attachmentServer server.AttachmentServer
	list, err := attachmentServer.List(topicIdInt, utils.UserIDUint(c.Ctx))
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(list)
}

// PostBuyBy 购买附件
func (c *Attachment) PostBuyBy(attachmentId string) utils.ResponseData {
	if authMiddleware, data := middleware.AuthMiddleware(c.Ctx); !authMiddleware {
		return *data
	}
	userId := utils.UserIDUint(c.Ctx)
	var attachmentServer server.AttachmentServer
	err := attachmentServer.Buy(db.GetIntID(attachmentId), userId)
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(nil)
}
