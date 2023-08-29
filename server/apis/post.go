package apis

import (
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
)

type Post struct {
	Ctx *gin.Context
}

func (p *Post) Post() utils.ResponseData {

	return utils.JsonSuccess(nil)
}
