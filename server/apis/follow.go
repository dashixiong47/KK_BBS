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
	if err != nil {
		return utils.JsonParameterError(err)
	}

	return utils.JsonSuccess(nil)
}

// GetBy 获取关注列表
func (f *Follow) GetBy(userId string) utils.ResponseData {
	// 验证用户是否已通过身份验证
	if authMiddleware, data := middleware.AuthMiddleware(f.Ctx); !authMiddleware {
		return *data
	}
	// 从上下文中获取用户ID
	userIdInt, _ := utils.DecryptID(userId)

	var paging utils.Paging
	paging.GetPaging(f.Ctx)
	var followServer server.FollowServer
	followList, err := followServer.GetFollowList(userIdInt, paging)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	return utils.JsonSuccess(followList)
}

// GetFansBy 获取fans列表
func (f *Follow) GetFansBy(userId string) utils.ResponseData {
	// 验证用户是否已通过身份验证
	if authMiddleware, data := middleware.AuthMiddleware(f.Ctx); !authMiddleware {
		return *data
	}
	// 从上下文中获取用户ID
	userIdInt, _ := utils.DecryptID(userId)
	var paging utils.Paging
	paging.GetPaging(f.Ctx)
	var followServer server.FollowServer
	followList, err := followServer.GetFansList(userIdInt, paging)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	return utils.JsonSuccess(followList)
}
