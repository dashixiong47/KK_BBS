package apis

import (
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
)

type User struct {
	Ctx *gin.Context
}

func (u *User) Get() utils.ResponseData {
	value, _ := u.Ctx.Get("id")
	var userInfo server.UserServer
	info, err := userInfo.GetUserInfo(int(value.(float64)))
	if err != nil {
		return utils.JsonParameterError("参数错误")
	}
	return utils.JsonSuccess(info)
}
func (u *User) GetBy(id string) utils.ResponseData {
	intID, err := utils.DecryptID(id)
	if err != nil {
		return utils.JsonParameterError("参数错误")
	}
	var userInfo server.UserServer
	info, err := userInfo.GetUserInfo(intID)
	if err != nil {
		return utils.JsonParameterError("参数错误")
	}
	return utils.JsonSuccess(info)
}
