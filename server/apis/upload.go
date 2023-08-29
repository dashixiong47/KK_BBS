package apis

import (
	"fmt"
	"github.com/dashixiong47/KK_BBS/server"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/gin-gonic/gin"
	"mime"
	"path/filepath"
	"time"
)

type Upload struct {
	Ctx *gin.Context
}

func (u *Upload) Post() utils.ResponseData {
	var uploadServer server.UploadServer
	// 获取文件
	file, _ := u.Ctx.FormFile("file")
	//判断文件大小 超出限制
	if file.Size > 1024*1024*10 {
		return utils.JsonFail("文件大小超出限制")
	}
	// 打开文件流
	fileReader, err := file.Open()
	if err != nil {
		return utils.JsonFail("打开文件失败")
	}
	defer fileReader.Close()
	// 计算 MD5
	md5Hash := utils.StreamToMD5(fileReader)
	query := uploadServer.Query(md5Hash)
	if query != "" {
		return utils.JsonSuccess(query)
	}
	// 获取文件路径
	path, _type := getPath(file.Filename, md5Hash)

	// 保存文件
	if err := u.Ctx.SaveUploadedFile(file, path); err != nil {
		return utils.JsonFail("保存文件失败")
	}
	// 保存文件信息到数据库
	uploadServer.Save(md5Hash, file.Filename, _type, path[1:])
	return utils.JsonSuccess(map[string]string{
		"url":  path[1:],
		"name": file.Filename,
	})
}

func getPath(filename, md5Hash string) (string, string) {
	//检查文件类型
	ext := filepath.Ext(filename)
	name := getFileName(md5Hash + ext)
	mimeType := mime.TypeByExtension(ext)
	var dst string
	switch mimeType {
	case "image/jpeg", "image/png":
		// 保存图片到 ./images 目录
		dst = fmt.Sprintf("./files/images/%s", name)
	case "video/mp4":
		// 保存视频到 ./videos 目录
		dst = fmt.Sprintf("./files/videos/%s", name)
	default:
		// 保存其他文件到 ./files 目录
		dst = fmt.Sprintf("./files/files/%s", name)
	}
	_type := ""
	if ext != "" {
		_type = ext[1:]
	}
	return dst, _type
}

// 根据时间生成文件名 2023/01/01/xxxx.jpg
func getFileName(filename string) string {
	return fmt.Sprintf("%s%s", time.Now().Format("2006/01/02/"), filename)
}
