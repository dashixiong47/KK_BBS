package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseData 定义了统一的响应结构
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Respond 用于发送统一格式的响应
func Respond(c *gin.Context, code int, message string, data interface{}) {
	response := ResponseData{
		Code:    code,
		Message: message,
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}

// JsonSuccess 成功响应
func JsonSuccess(data interface{}) ResponseData {
	return ResponseData{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

// JsonError 失败响应
func JsonError(code int, message string) ResponseData {
	return ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

// JsonParameterError 参数错误
func JsonParameterError(message ...string) ResponseData {
	defaultMessage := "参数错误"
	var responseMessage string

	if len(message) > 0 && message[0] != "" {
		responseMessage = message[0]
	} else {
		responseMessage = defaultMessage
	}
	return ResponseData{
		Code:    400,
		Message: responseMessage,
		Data:    nil,
	}
}

// JsonFail 失败响应
func JsonFail(message ...string) ResponseData {
	defaultMessage := "操作失败"
	var responseMessage string

	if len(message) > 0 && message[0] != "" {
		responseMessage = message[0]
	} else {
		responseMessage = defaultMessage
	}

	return ResponseData{
		Code:    500,
		Message: responseMessage,
		Data:    nil,
	}
}
