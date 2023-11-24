package admin

import (
	"errors"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/klog"
)

type UserServer struct {
}

// GetUserList 获取用户信息
func (u *UserServer) GetUserList(paging utils.Paging) (interface{}, error) {
	var docs []models.User
	var count int64
	// 从数据库中查询
	tx := paging.SetPaging(db.DB).
		Model(&models.User{}).
		Order("id desc")
	err := tx.Count(&count).Error
	if err != nil {
		klog.Error("GetUserInfo", err)
		return nil, errors.New("unknown")
	}
	if err := tx.Find(&docs).Error; err != nil {
		klog.Error("GetUserInfo", err)
		return nil, errors.New("unknown")
	}

	return map[string]any{
		"list":  docs,
		"total": count,
	}, nil
}

// GetUserInfo 获取用户信息
func (u *UserServer) GetUserInfo(id int) (any, error) {
	var userInfo models.SysUser
	// 如果缓存中没有，再从数据库中获取
	if err := db.DB.Where("id = ?", id).First(&userInfo).Error; err != nil {

		return nil, errors.New("user_not_found")
	}
	return userInfo, nil
}
