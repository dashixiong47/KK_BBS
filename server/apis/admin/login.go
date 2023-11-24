package admin

import (
	"github.com/dashixiong47/KK_BBS/server/admin"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/captcha"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Ctx *gin.Context
}
type sysUser struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Code      string `json:"code"`
	CaptchaId string `json:"captchaId"`
}

// Post 登录
func (c *Login) Post() utils.ResponseData {
	var info sysUser
	err := c.Ctx.ShouldBindJSON(&info)
	if err != nil {
		return utils.JsonFail("json_error")
	}

	var loginServer admin.LoginServer
	if info.CaptchaId == "" {
		return utils.JsonFail("captcha_id_is_empty")
	}
	if info.Code == "" {
		return utils.JsonFail("code_is_empty")
	}
	verifyCaptcha := captcha.VerifyCaptcha(info.CaptchaId, info.Code)
	if !verifyCaptcha {
		return utils.JsonFail("code_error")
	}
	if info.Password == "" {
		return utils.JsonFail("password_error")
	}
	md5 := utils.MD5(info.Password)
	userInfo, err := loginServer.Login(info.Username, md5)
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(userInfo)
}
