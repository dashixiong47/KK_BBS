package middleware

import (
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/jwt"
	"github.com/gin-gonic/gin"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

type CustomClaims struct {
	User string `json:"user"`
	ID   string `json:"id"`
	// 其他字段
}

var whiteList = []string{
	"/api/v1/topic/list",
	"/api/v1/topic/",
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		latency := time.Since(startTime)
		log.Printf("请求方法: %s, 请求路径: %s, 请求耗时: %s", c.Request.Method, c.Request.URL.Path, latency)
	}
}

//func AuthMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		header := c.GetHeader("Authorization")
//		// 白名单
//		for _, path := range whiteList {
//			if path == c.Request.URL.Path || strings.HasPrefix(c.Request.URL.Path, path) {
//				c.Next()
//				return
//			}
//		}
//		if header == "" {
//			c.JSON(200, utils.JsonError(401, "please_login"))
//			c.Abort()
//			return
//		} else {
//			user, err := jwt.ParseToken(header)
//			if err != nil {
//				c.JSON(200, utils.JsonError(401, "token_error"))
//				c.Abort()
//				return
//			}
//			c.Set("user", user.Claims.(jwtv5.MapClaims)["user"])
//			c.Set("id", user.Claims.(jwtv5.MapClaims)["id"])
//		}
//		c.Next()
//	}
//}

func AuthMiddleware(c *gin.Context) (bool, *utils.ResponseData) {
	header := c.GetHeader("Authorization")

	var jsonError utils.ResponseData
	if header == "" {
		jsonError = utils.JsonError(401, "please_login")
		return false, &jsonError
	} else {
		user, err := jwt.ParseToken(header)
		if err != nil {
			jsonError = utils.JsonError(401, "token_error")
			return false, &jsonError
		}
		c.Set("user", user.Claims.(jwtv5.MapClaims)["user"])
		c.Set("id", user.Claims.(jwtv5.MapClaims)["id"])
		return true, nil
	}
}
