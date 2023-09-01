package router

import (
	"fmt"
	"github.com/dashixiong47/KK_BBS/apis"
	"github.com/dashixiong47/KK_BBS/config"
	"github.com/dashixiong47/KK_BBS/middleware"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"github.com/gin-gonic/gin"
)

func init() {
	r := gin.Default() // 创建一个 Gin 实例
	r.Static("/files", "./files")

	if config.SettingsConfig.Application.Mode == "dev" {
		r.Use(crossAllow)
	}

	r.Use(middleware.Logger())
	port := fmt.Sprintf(":%d", config.SettingsConfig.Application.Port)
	klog.Info("Server starting on port %s", port)
	registerRoutes(r) // 注册路由
	err := r.Run(port)
	if err != nil {
		klog.Error("Server failed to start: %v", err)
	}
}
func registerRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware())
	RegisterRoutes(v1, &apis.User{})
	RegisterRoutes(v1, &apis.Upload{})
	RegisterRoutes(v1, &apis.Post{})

	// 不需要认证的路由
	noCheck := r.Group("/api/v1")
	RegisterRoutes(noCheck, &apis.Captcha{})
	RegisterRoutes(noCheck, &apis.Login{})
	RegisterRoutes(noCheck, &apis.Group{})

}

func crossAllow(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
