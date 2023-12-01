package admin

import (
	"errors"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/server/admin"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
)

type Group struct {
	Ctx *gin.Context
}

// Get 获取
func (c *Group) Get() utils.ResponseData {
	var groupServer admin.GroupServer
	var paging utils.Paging
	paging.GetPaging(c.Ctx)
	name := c.Ctx.DefaultQuery("name", "")
	info, err := groupServer.GroupInfo(paging, name)
	if err != nil {
		return utils.JsonFail()
	}
	return utils.JsonSuccess(info)
}

// Post 添加
func (c *Group) Post() utils.ResponseData {
	var doc models.Group
	var groupServer admin.GroupServer
	if err := c.Ctx.ShouldBind(&doc); err != nil {
		return utils.JsonParameterError(errors.New("json_error"))
	}
	info, err := groupServer.GroupAdd(doc)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	return utils.JsonSuccess(info)
}

// Put 修改
func (c *Group) Put() utils.ResponseData {
	var doc models.Group
	var groupServer admin.GroupServer
	if err := c.Ctx.ShouldBind(&doc); err != nil {
		return utils.JsonParameterError(errors.New("json_error"))
	}
	info, err := groupServer.GroupUpdate(doc)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	return utils.JsonSuccess(info)
}

// Delete 删除
func (c *Group) Delete() utils.ResponseData {
	var groupServer admin.GroupServer
	id := c.Ctx.DefaultQuery("id", "")
	info, err := groupServer.GroupDelete(id)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	return utils.JsonSuccess(info)
}
