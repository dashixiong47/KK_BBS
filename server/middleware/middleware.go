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

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		latency := time.Since(startTime)
		log.Printf("请求方法: %s, 请求路径: %s, 请求耗时: %s", c.Request.Method, c.Request.URL.Path, latency)
	}
}
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(401, utils.JsonError(401, "请登录后访问"))
			c.Abort()
			return
		} else {
			user, err := jwt.ParseToken(header)
			c.Set("user", user.Claims.(jwtv5.MapClaims)["user"])
			c.Set("id", user.Claims.(jwtv5.MapClaims)["id"])
			if err != nil {
				c.JSON(401, utils.JsonError(401, "请登录后访问"))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
