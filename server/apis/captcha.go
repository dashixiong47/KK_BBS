package apis

import (
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/captcha"
	"github.com/dashixiong47/KK_BBS/utils/stmp"
	"github.com/gin-gonic/gin"
)

type Captcha struct {
	Ctx *gin.Context
}
type fromData struct {
	Email     string `json:"email" binding:"required"`
	Code      string `json:"code" binding:"required"`
	CaptchaId string `json:"captchaId" binding:"required"`
}

func (c *Captcha) Get() utils.ResponseData {
	id, s, err := captcha.GenerateCaptcha()
	if err != nil {
		return utils.JsonFail("生成验证码失败")

	}
	return utils.JsonSuccess(gin.H{
		"captchaId": id,
		"captcha":   s,
	})
}

// PostCode 发送验证码
func (c *Captcha) PostCode() utils.ResponseData {
	var info fromData
	err := c.Ctx.ShouldBindJSON(&info)
	if err != nil {
		return utils.JsonFail("json_error")
	}
	// 发送验证码
	err = stmp.SendVerificationEmail(info.Email, info.Email)
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(nil)
}
