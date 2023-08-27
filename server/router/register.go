package router

import (
	"fmt"
	"github.com/dashixiong47/KK_BBS/apis"
	"github.com/dashixiong47/KK_BBS/config"
	"github.com/dashixiong47/KK_BBS/middleware"
	"github.com/dashixiong47/KK_BBS/utils/print"
	"github.com/gin-gonic/gin"
)

func init() {
	r := gin.Default() // 创建一个 Gin 实例
	r.Use(middleware.Logger())
	port := fmt.Sprintf(":%d", config.SettingsConfig.Application.Port)
	print.Info("Server starting on port %s", port)
	registerRoutes(r) // 注册路由
	err := r.Run(port)
	if err != nil {
		print.Error("Server failed to start: %v", err)
	}
}
func registerRoutes(r *gin.Engine) {
	v1 := r.Group("/v1/api")
	v1.Use(middleware.AuthMiddleware())
	// 不需要认证的路由
	noCheck := r.Group("/v1/api")
	RegisterRoutes(noCheck, &apis.Login{})
	RegisterRoutes(noCheck, &apis.MyController{}) // 调用 RegisterRoutes 函数注册 MyController 的方法

}
