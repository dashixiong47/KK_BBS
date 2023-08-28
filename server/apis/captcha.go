package apis

import (
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/captcha"
	"github.com/gin-gonic/gin"
)

type Captcha struct {
	Ctx *gin.Context
}

func (c *Captcha) Get() utils.ResponseData {
	id, s, err := captcha.GenerateCaptcha(c.Ctx)
	if err != nil {
		return utils.JsonFail("生成验证码失败")

	}
	return utils.JsonSuccess(gin.H{
		"captchaId": id,
		"captcha":   s,
	})
}

func (c *Captcha) Post() utils.ResponseData {
	verify := captcha.VerifyCaptcha(c.Ctx)
	if !verify {
		return utils.JsonParameterError("验证码错误")
	}
	return utils.JsonSuccess(nil)
}
