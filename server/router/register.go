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

	//// 全局使用 CORS 中间件
	//_config := cors.DefaultConfig()
	//_config.AllowAllOrigins = true // 允许所有域名访问
	//_config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	//_config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	//
	//r.Use(cors.New(_config))

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

	// 不需要认证的路由
	noCheck := r.Group("/api/v1")
	RegisterRoutes(noCheck, &apis.Captcha{})
	RegisterRoutes(noCheck, &apis.Login{})
	RegisterRoutes(noCheck, &apis.Group{})

}
