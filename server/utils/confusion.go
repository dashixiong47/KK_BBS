package utils

import (
	"encoding/base64"
	"errors"
	"github.com/dashixiong47/KK_BBS/config"
	"strconv"
)

var key = config.SettingsConfig.Application.IdConfusion

// EncryptID 加密整数 ID
func EncryptID(originalID int) string {

	// 使用 XOR 加密
	encryptedID := originalID ^ key
	// 将加密后的整数转换为字节串
	encryptedBytes := []byte(strconv.Itoa(encryptedID))

	// 使用 Base64 编码
	return base64.StdEncoding.EncodeToString(encryptedBytes)
}

// DecryptID 解密加密后的字符串，返回原始 ID
func DecryptID(encodedString string) (int, error) {
	// 使用 Base64 解码
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return 0, errors.New("invalid_id")
	}

	// 将字节串转换回整数
	decodedID, err := strconv.Atoi(string(decodedBytes))
	if err != nil {
		return 0, errors.New("invalid_id")
	}

	// 使用 XOR 解密
	return decodedID ^ key, nil
}

// 从redis里获取
