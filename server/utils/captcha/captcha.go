package captcha

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type InputCaptcha struct {
	CaptchaId string `json:"captchaId"`
	Value     string `json:"value"`
}

var store = base64Captcha.DefaultMemStore

func GenerateCaptcha(c *gin.Context) (captchaId string, b64s string, err error) {
	// 创建图片验证码配置
	driver := base64Captcha.DriverString{
		Height:          80,
		Width:           240,
		NoiseCount:      0,
		ShowLineOptions: 0,
		Length:          4,
		Source:          "1234567890",
	}
	captcha := base64Captcha.NewCaptcha(&driver, store)

	// 生成验证码
	id, b64s, err := captcha.Generate()
	if err != nil {
		return "", "", err
	}
	return id, b64s, nil

}
func VerifyCaptcha(c *gin.Context) bool {
	var Captcha InputCaptcha
	_ = c.ShouldBindJSON(&Captcha)
	return store.Verify(Captcha.CaptchaId, Captcha.Value, true)
}
