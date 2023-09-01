package apis

import (
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
)

type Post struct {
	Ctx *gin.Context
}

func (p *Post) Post() utils.ResponseData {
	var doc models.Post
	err := p.Ctx.ShouldBindJSON(&doc)
	if err != nil {
		return utils.JsonParameterError(err.Error())
	}
	if err := db.DB.Create(&doc).Error; err != nil {
		return utils.JsonFail(err.Error())
	}

	return utils.JsonSuccess(doc)
}
