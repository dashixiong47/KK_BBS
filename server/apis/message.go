package apis

import (
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/middleware"
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
)

type Message struct {
	Ctx *gin.Context
}

func (c *Message) Get() utils.ResponseData {
	if authMiddleware, data := middleware.AuthMiddleware(c.Ctx); !authMiddleware {
		return *data
	}
	userId := utils.UserIDInt(c.Ctx)
	var paging utils.Paging
	paging.GetPaging(c.Ctx)
	var userInfo server.MessageServer
	info, err := userInfo.MessageInfo(c.Ctx, userId, paging)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	return utils.JsonSuccess(info)
}

// GetUnread 获取未读消息数量
func (c *Message) GetUnread() utils.ResponseData {
	if authMiddleware, data := middleware.AuthMiddleware(c.Ctx); !authMiddleware {
		return *data
	}
	userId := utils.UserIDInt(c.Ctx)
	var userInfo server.MessageServer
	info, err := userInfo.UnreadMessageInfo(userId)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	return utils.JsonSuccess(info)
}

// PostReadBy 单个已读
func (c *Message) PostReadBy(id string) utils.ResponseData {

	if authMiddleware, data := middleware.AuthMiddleware(c.Ctx); !authMiddleware {
		return *data
	}
	userId := utils.UserIDUint(c.Ctx)
	var userInfo server.MessageServer
	err := userInfo.ReadBy(uint(db.GetIntID(id)), userId)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	return utils.JsonSuccess(nil)
}

// PostReadAll 全部已读
func (c *Message) PostReadAll() utils.ResponseData {
	if authMiddleware, data := middleware.AuthMiddleware(c.Ctx); !authMiddleware {
		return *data
	}
	userId := utils.UserIDInt(c.Ctx)

	var userInfo server.MessageServer
	err := userInfo.ReadAll(userId)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	return utils.JsonSuccess(nil)
}
