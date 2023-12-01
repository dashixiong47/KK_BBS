package apis

import (
	"github.com/dashixiong47/KK_BBS/middleware"
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Ctx *gin.Context
}
type userInfo struct {
	Avatar       string `json:"avatar"`
	Nickname     string `json:"nickname"`
	Background   string `json:"background"`
	Introduction string `json:"introduction"`
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

// Put 修改用户信息
func (u *User) Put() utils.ResponseData {
	if authMiddleware, data := middleware.AuthMiddleware(u.Ctx); !authMiddleware {
		return *data
	}
	value, _ := u.Ctx.Get("id")
	var user userInfo
	err := u.Ctx.ShouldBind(&user)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	log.Println(user)
	var userServer server.UserServer
	info, err := userServer.UpdateUserInfo(int(value.(float64)), user.Nickname, user.Avatar, user.Background, user.Introduction)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	return utils.JsonSuccess(info)
}
