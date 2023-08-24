package apis

import (
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
)

type Login struct {
	Ctx *gin.Context
}

func (c *Login) Post() {
	var info map[string]string
	err := c.Ctx.ShouldBindJSON(&info)
	if err != nil {
		utils.JsonFail(c.Ctx, "参数错误")
		return
	}

	utils.JsonSuccess(c.Ctx, info)
}
