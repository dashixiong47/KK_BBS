package apis

import (
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Ctx *gin.Context
}
type user struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Post 登录
func (c *Login) Post() utils.ResponseData {
	var info user
	err := c.Ctx.ShouldBindJSON(&info)
	if err != nil {
		return utils.JsonFail("参数错误")
	}
	md5 := utils.MD5(info.Password)
	var loginServer server.LoginServer
	userInfo, err := loginServer.Login(info.Username, md5)
	if err != nil {
		return utils.JsonFail("用户名或密码错误")
	}
	return utils.JsonSuccess(userInfo)
}
