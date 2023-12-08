package apis

import (
	"github.com/dashixiong47/KK_BBS/middleware"
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
)

type Follow struct {
	Ctx *gin.Context
}

// PostBy 用于创建关注
func (f *Follow) PostBy(userId string) utils.ResponseData {
	// 验证用户是否已通过身份验证
	if authMiddleware, data := middleware.AuthMiddleware(f.Ctx); !authMiddleware {
		return *data
	}

	// 从上下文中获取用户ID
	idInt := utils.UserIDInt(f.Ctx)

	// 获取关注用户ID
	followIdInt, err := utils.DecryptID(userId)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	var followServer server.FollowServer
	err = followServer.CreateFollow(idInt, followIdInt)

	return utils.JsonSuccess(nil)
}
