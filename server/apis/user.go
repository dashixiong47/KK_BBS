package apis

import (
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Ctx *gin.Context
}

func (u *User) GetBy(id string) utils.ResponseData {
	intID, err := utils.DecryptID(id)
	if err != nil {
		return utils.JsonParameterError("参数错误")
	}
	log.Println(intID)
	var userInfo server.UserServer
	info, err := userInfo.UserInfo(intID)
	if err != nil {
		return utils.JsonParameterError("参数错误")
	}
	return utils.JsonSuccess(info)
}
