package apis

import (
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Search struct {
	Ctx *gin.Context
}

func (c *Search) Get() utils.ResponseData {
	q := c.Ctx.DefaultQuery("q", "")
	typeStr := c.Ctx.DefaultQuery("type", "-1")
	groupStr := c.Ctx.DefaultQuery("group", "-1")
	typeInt, err := strconv.Atoi(typeStr)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	groupInt, err := strconv.Atoi(groupStr)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	var searchServer server.SearchServer
	search, err := searchServer.Search(q, 0, 10, typeInt, groupInt)
	if err != nil {
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(search)

}
