package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// 计算字符串的 MD5 值
func MD5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	hashString := hex.EncodeToString(hash.Sum(nil))
	return hashString
}

// 计算 io.Reader 流的 MD5 哈希
func StreamToMD5(stream io.Reader) string {
	hash := md5.New()
	io.Copy(hash, stream)
	return hex.EncodeToString(hash.Sum(nil))
}
