package upload

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"github.com/gin-gonic/gin"
	"log"
	"mime"
	"path/filepath"
	"time"
)

func UploadFile(c *gin.Context) (any, error) {
	// 获取文件
	file, err := c.FormFile("file")
	if err != nil {
		return nil, err
	}
	//判断文件大小 超出限制
	if file.Size > 1024*1024*10 {
		return nil, errors.New("file_size_exceeds_limit")
	}
	// 打开文件流
	fileReader, err := file.Open()
	if err != nil {
		return nil, errors.New("failed_to_open_file")
	}
	defer fileReader.Close()
	// 计算 MD5
	md5Hash := utils.StreamToMD5(fileReader)
	query := fileQuery(md5Hash)
	if query != "" {
		return query, nil
	}
	// 获取文件路径
	path, _type := getPath(file.Filename, md5Hash)

	// 保存文件
	if err := c.SaveUploadedFile(file, path); err != nil {
		return nil, errors.New("save_file_info_error")
	}
	// 保存文件信息到数据库
	save, err := fileSave(md5Hash, file.Filename, _type, path[1:])
	if err != nil {
		return nil, errors.New("save_file_info_error")
	}

	return map[string]interface{}{
		"id":   save.ID,
		"url":  path[1:],
		"name": save.FileName,
	}, nil
}

func fileQuery(md5 string) any {

	var fileInfo models.File
	ctx := context.Background()
	// 从redis中查询
	if result, err := db.Rdb.Get(ctx, fmt.Sprintf("file:%v", md5)).Result(); err == nil {
		log.Println("从redis中查询")
		var data map[string]interface{}
		_ = json.Unmarshal([]byte(result), &data)
		return data
	}
	log.Println("从数据库中查询")
	// 从数据库中查询
	err := db.DB.Where("md5 = ?", md5).First(&fileInfo).Error
	if err != nil {
		return ""
	}
	_ = saveToRedis(fileInfo)
	return map[string]interface{}{
		"id":   fileInfo.ID,
		"url":  fileInfo.Url,
		"name": fileInfo.FileName,
	}
}
func fileSave(md5, filename, _type, path string) (models.File, error) {
	var fileInfo models.File
	fileInfo.Md5 = md5
	fileInfo.FileName = filename
	fileInfo.Type = _type
	fileInfo.Url = path

	err := db.DB.Create(&fileInfo).Error
	if err != nil {
		klog.Error(err.Error())
		return models.File{}, err
	}
	_ = saveToRedis(fileInfo)
	return fileInfo, nil
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

// 存到redis
func saveToRedis(file models.File) error {
	// ExpireTime 过期时间
	const ExpireTime = 2 * time.Hour
	data := map[string]interface{}{
		"id":   file.ID,
		"name": file.FileName,
		"url":  file.Url,
	}
	marshal, _ := json.Marshal(data)
	// 保存到redis 过期时间2小时
	ctx := context.Background()
	err := db.Rdb.Set(ctx, fmt.Sprintf("file:%v", file.Md5), marshal, ExpireTime).Err()
	if err != nil {
		return err
	}
	return nil
}
