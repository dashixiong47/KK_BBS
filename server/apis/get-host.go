package apis

import (
	"github.com/dashixiong47/KK_BBS/config"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
)

type Host struct {
	Ctx *gin.Context
}

// Get 获取主机信息
func (h *Host) Get() utils.ResponseData {
	return utils.JsonSuccess(config.SettingsConfig.Application.FileHost)
}
