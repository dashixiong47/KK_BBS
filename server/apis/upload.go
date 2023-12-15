package apis

import (
	"github.com/dashixiong47/KK_BBS/middleware"
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
)

type Upload struct {
	Ctx *gin.Context
}

func (u *Upload) Post() utils.ResponseData {
	if authMiddleware, data := middleware.AuthMiddleware(u.Ctx); !authMiddleware {
		return *data
	}
	var upload server.UploadServer
	info, err := upload.UploadFile(u.Ctx)
	if err != nil {
		return utils.JsonParameterError(err)
	}
	return utils.JsonSuccess(info)
}

//func (u *Upload) PostVideo() utils.ResponseData {
//	if authMiddleware, data := middleware.AuthMiddleware(u.Ctx); !authMiddleware {
//		return *data
//	}
//	var upload server.UploadServer
//	info, err := upload.UploadVideo(u.Ctx)
//	if err != nil {
//		return utils.JsonFail(err)
//	}
//	return utils.JsonSuccess(info)
//}
