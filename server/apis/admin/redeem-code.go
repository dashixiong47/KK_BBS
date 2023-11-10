package admin

import (
	"github.com/dashixiong47/KK_BBS/server/admin"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"github.com/gin-gonic/gin"
)

type RedeemCode struct {
	Ctx *gin.Context
}
type formData struct {
	Quantity int `json:"quantity"`
	Type     int `json:"type"`
	Number   int `json:"number"`
}

// PostCreate 生成兑换码
func (r *RedeemCode) PostCreate() utils.ResponseData {
	var formData formData
	err := r.Ctx.ShouldBindJSON(&formData)
	if err != nil {
		klog.Error("RedeemCode PostCreate", err)
		return utils.JsonFail(err)
	}
	var redeemCodeServer admin.RedeemCodeServer
	docs, err := redeemCodeServer.Create(formData.Quantity, formData.Type, formData.Number)
	if err != nil {
		klog.Error("RedeemCode PostCreate", err)
		return utils.JsonFail(err)
	}
	return utils.JsonSuccess(docs)
}
