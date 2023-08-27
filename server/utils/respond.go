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
func JsonParameterError(message string) ResponseData {
	return ResponseData{
		Code:    400,
		Message: message,
		Data:    nil,
	}
}

// JsonFail 失败响应
func JsonFail(message string) ResponseData {
	return ResponseData{
		Code:    500,
		Message: message,
		Data:    nil,
	}
}
