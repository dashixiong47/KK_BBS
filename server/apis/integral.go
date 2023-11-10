package apis

import (
	"github.com/dashixiong47/KK_BBS/middleware"
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"github.com/gin-gonic/gin"
)

type Integral struct {
	Ctx *gin.Context
}

// PostRecharge 充值积分
func (c *Integral) PostRecharge() utils.ResponseData {
	if authMiddleware, data := middleware.AuthMiddleware(c.Ctx); !authMiddleware {
		return *data
	}

	var integral server.IntegralServer
	err := integral.Recharge(utils.UserIDUint(c.Ctx))
	if err != nil {
		klog.Error("integral recharge error: %v", err)
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(nil)
}
