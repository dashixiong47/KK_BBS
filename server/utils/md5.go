package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 计算字符串的 MD5 值
func MD5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString
}
