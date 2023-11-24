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
	Username  string `json:"username"`
	Password  string `json:"password"`
	Code      string `json:"code"`
	Email     string `json:"email"`
	VCode     string `json:"vCode"`
	LoginType *int   `json:"loginType"`
	CaptchaId string `json:"captchaId"`
}
type register struct {
	Email     string `json:"email" binding:"required"`
	Code      string `json:"code" binding:"required"`
	VCode     string `json:"vCode" binding:"required"`
	CaptchaId string `json:"captchaId" binding:"required"`
}

// Post 登录
func (c *Login) Post() utils.ResponseData {
	var info user
	err := c.Ctx.ShouldBindJSON(&info)
	if err != nil {
		return utils.JsonFail("json_error")
	}

	var loginServer server.LoginServer
	if info.LoginType == nil || *info.LoginType == 1 {
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
	} else {
		if info.VCode == "" {
			return utils.JsonFail("vcode_is_empty")
		}
		if info.Email == "" {
			return utils.JsonFail("email_is_empty")
		}
		userInfo, err := loginServer.Register(info.Email, info.VCode)
		if err != nil {
			return utils.JsonFail(err)
		}
		return utils.JsonSuccess(userInfo)
	}

}

// PostRegister 注册
func (c *Login) PostRegister() utils.ResponseData {
	var info register
	err := c.Ctx.ShouldBindJSON(&info)
	//if err != nil {
	//	return utils.JsonFail("json_error")
	//}
	//verifyCaptcha := captcha.VerifyCaptcha(info.CaptchaId, info.Code)
	//if !verifyCaptcha {
	//	return utils.JsonFail("code_error")
	//}
	var loginServer server.LoginServer
	userInfo, err := loginServer.Register(info.Email, info.VCode)
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(userInfo)
}
