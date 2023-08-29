package apis

import (
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/captcha"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Ctx *gin.Context
}
type user struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Code      string `json:"code" binding:"required"`
	LoginType *int   `json:"loginType"`
	CaptchaId string `json:"captchaId"`
}

// Post 登录
func (c *Login) Post() utils.ResponseData {
	var info user
	err := c.Ctx.ShouldBindJSON(&info)
	if err != nil {
		return utils.JsonFail(err.Error())
	}
	verifyCaptcha := captcha.VerifyCaptcha(info.CaptchaId, info.Code)
	if !verifyCaptcha {
		return utils.JsonFail("验证码错误")
	}
	md5 := utils.MD5(info.Password)
	var loginServer server.LoginServer
	userInfo, err := loginServer.Login(info.Username, md5)
	if err != nil {
		return utils.JsonFail("用户名或密码错误")
	}
	return utils.JsonSuccess(userInfo)
}
