package admin

import (
	"github.com/dashixiong47/KK_BBS/middleware"
	"github.com/dashixiong47/KK_BBS/server/admin"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
)

type User struct {
	Ctx *gin.Context
}

// GetList 获取用户信息
func (c *User) GetList() utils.ResponseData {
	// 验证用户是否已通过身份验证
	if authMiddleware, data := middleware.AuthMiddleware(c.Ctx); !authMiddleware {
		return *data
	}
	var paging utils.Paging
	paging.GetPaging(c.Ctx)
	var userServer admin.UserServer
	info, err := userServer.GetUserList(paging)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	return utils.JsonSuccess(info)

}

// Get 获取用户信息
func (u *User) Get() utils.ResponseData {
	if authMiddleware, data := middleware.AuthMiddleware(u.Ctx); !authMiddleware {
		return *data
	}
	value, _ := u.Ctx.Get("id")
	var userInfo admin.UserServer
	info, err := userInfo.GetUserInfo(int(value.(float64)))
	if err != nil {
		return utils.JsonParameterError(err)
	}
	return utils.JsonSuccess(info)
}
