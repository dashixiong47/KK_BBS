package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dashixiong47/KK_BBS/config"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"log"
	"time"
)

type UploadServer struct{}

func (u *UploadServer) Query(md5 string) any {
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
		"url":  config.SettingsConfig.Application.Host + fileInfo.Url,
		"name": fileInfo.FileName,
	}
}

func (u *UploadServer) Save(md5, filename, _type, path string) (models.File, error) {
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

// 存到redis
func saveToRedis(file models.File) error {
	// ExpireTime 过期时间
	const ExpireTime = 2 * time.Hour
	data := map[string]interface{}{
		"id":   file.ID,
		"name": file.FileName,
		"url":  config.SettingsConfig.Application.Host + file.Url,
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
