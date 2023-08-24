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

// 成功响应
func JsonSuccess(g *gin.Context, data interface{}) {
	Respond(g, 200, "success", data)
}

// 失败响应
func JsonFail(g *gin.Context, message string) {
	Respond(g, 500, message, nil)
}
