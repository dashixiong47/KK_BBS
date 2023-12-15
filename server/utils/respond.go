package utils

import (
	"github.com/dashixiong47/KK_BBS/config"
	"github.com/dashixiong47/KK_BBS/locales"
)

// ResponseData 定义了统一的响应结构
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// JsonSuccess 成功响应
func JsonSuccess(data interface{}) ResponseData {
	return ResponseData{
		Code:    200,
		Message: getMessage("success"),
		Data:    data,
	}
}

// JsonError 失败响应
func JsonError(code int, message any) ResponseData {
	defaultMessage := "unknown_error"
	switch message.(type) {
	case string:
		defaultMessage = message.(string)
	case error:
		defaultMessage = message.(error).Error()
	}
	return ResponseData{
		Code:    code,
		Message: getMessage(defaultMessage),
		Data:    nil,
	}
}

// JsonParameterError 参数错误
func JsonParameterError(message ...any) ResponseData {
	defaultMessage := "parameter_error"

	if len(message) > 0 && message[0] != "" {
		switch message[0].(type) {
		case string:
			defaultMessage = message[0].(string)
		case error:
			defaultMessage = message[0].(error).Error()
		}
	}
	return ResponseData{
		Code:    400,
		Message: getMessage(defaultMessage),
		Data:    nil,
	}
}

// JsonFail 失败响应
func JsonFail(message ...any) ResponseData {
	defaultMessage := "parameter_error"

	if len(message) > 0 && message[0] != "" {
		switch message[0].(type) {
		case string:
			defaultMessage = message[0].(string)
		case error:
			defaultMessage = message[0].(error).Error()
		}
	}

	return ResponseData{
		Code:    500,
		Message: getMessage(defaultMessage),
		Data:    nil,
	}
}

// 从i18里面取值
func getMessage(message string) string {
	// 获取语言配置
	lang := config.SettingsConfig.Application.I18
	// 获取该语言的所有消息
	messages, ok := locales.I18[lang]
	if !ok {
		return message // 语言不存在
	}
	// 获取特定消息
	msg, ok := messages[message]
	if !ok {
		return message // 消息不存在
	}
	return msg
}
