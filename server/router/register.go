package router

import (
	"fmt"
	"github.com/dashixiong47/KK_BBS/apis"
	"github.com/dashixiong47/KK_BBS/config"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	r := gin.Default() // 创建一个 Gin 实例
	port := fmt.Sprintf(":%d", config.SettingsConfig.Application.Port)
	utils.Info("Server starting on port %s", port)
	registerRoutes(r) // 注册路由
	err := r.Run(port)
	if err != nil {
		utils.Error("Server failed to start: %v", err)
	}
}
func registerRoutes(r *gin.Engine) {
	v1 := r.Group("/v1/api")
	RegisterRoutes(v1, &apis.MyController{}) // 调用 RegisterRoutes 函数注册 MyController 的方法
	// 不需要认证的路由

	noCheck := r.Group("/v1/api")
	RegisterRoutes(noCheck, &apis.Login{})

}
