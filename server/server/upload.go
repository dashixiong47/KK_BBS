package server

import (
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils/klog"
)

type UploadServer struct{}

func (u *UploadServer) Query(md5 string) any {
	var fileInfo models.File
	err := db.DB.Where("md5 = ?", md5).First(&fileInfo).Error
	if err != nil {
		return ""
	}
	return map[string]string{
		"url":  fileInfo.Url,
		"name": fileInfo.FileName,
	}
}

func (u *UploadServer) Save(md5, filename, _type, path string) {
	var fileInfo models.File
	fileInfo.Md5 = md5
	fileInfo.FileName = filename
	fileInfo.Type = _type
	fileInfo.Url = path
	err := db.DB.Create(&fileInfo).Error
	if err != nil {
		klog.Error(err.Error())
	}

}
