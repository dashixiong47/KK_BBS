package apis

import (
	"github.com/dashixiong47/KK_BBS/middleware"
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
)

type User struct {
	Ctx *gin.Context
}

// Get 获取用户信息
func (u *User) Get() utils.ResponseData {
	if authMiddleware, data := middleware.AuthMiddleware(u.Ctx); !authMiddleware {
		return *data
	}
	value, _ := u.Ctx.Get("id")
	var userInfo server.UserServer
	info, err := userInfo.GetUserInfo(int(value.(float64)))
	if err != nil {
		return utils.JsonParameterError(err)
	}
	return utils.JsonSuccess(info)
}

// GetBy 获取用户信息
func (u *User) GetBy(id string) utils.ResponseData {
	intID, err := utils.DecryptID(id)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	var userInfo server.UserServer
	info, err := userInfo.GetUserInfo(intID)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	return utils.JsonSuccess(info)
}
