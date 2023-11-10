package data

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/dashixiong47/KK_BBS/config"
)

// CreateRedeemCode 生成兑换码
func CreateRedeemCode() (string, error) {
	redeemKey := config.SettingsConfig.Application.RedeemKey
	// 生成一个16字节的随机数
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	// 将随机数编码为base64格式的字符串
	encodedRandom := base64.StdEncoding.EncodeToString(randomBytes)
	encodedString := base64.StdEncoding.EncodeToString([]byte(redeemKey + encodedRandom))
	return encodedString, nil
}
