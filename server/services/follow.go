package services

import (
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
)

// IsUser 是否存在该用户
func IsUser(userId int) bool {
	var count int64
	db.DB.Model(models.User{}).Where("id = ?", userId).
		Count(&count)
	return count > 0
}

// GetUserFollow 获取用户关注的用户
func GetUserFollow(userId uint) int64 {
	var count int64 = 0
	db.DB.Debug().Model(models.Follow{}).
		Where("user_id=?", userId).
		Count(&count)
	return count
}

// GetUserFans 获取粉丝
func GetUserFans(userId uint) int64 {
	var count int64 = 0
	db.DB.Debug().Model(models.Follow{}).
		Where("follow_id=?", userId).
		Count(&count)
	return count
}
