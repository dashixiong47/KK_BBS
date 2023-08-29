package apis

import (
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
)

type Group struct {
	Ctx *gin.Context
}

func (g *Group) Get() utils.ResponseData {
	var groupServer server.GroupServer
	info, err := groupServer.GroupInfo()
	if err != nil {
		return utils.JsonFail()
	}
	return utils.JsonSuccess(info)
}
