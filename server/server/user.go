package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/server/data"
	"time"
)

type UserServer struct {
}

// GetUserInfo 获取用户信息
func (u *UserServer) GetUserInfo(id int) (any, error) {
	var userInfo models.User
	// 先从缓存中获取
	info := getRedisUserInfo(id)
	// 判断info是否有值
	if len(info) > 0 {
		return info, nil
	}
	// 如果缓存中没有，再从数据库中获取
	if err := db.DB.Where("id = ?", id).First(&userInfo).Error; err != nil {

		return nil, errors.New("user_not_found")
	}
	userInfo.Coins = data.UserIntegral(userInfo.ID)
	// 将数据存到缓存中
	setRedisUserInfo(userInfo)

	return userInfo, nil
}

// GetUserDetail 获取用户详情
func getRedisUserInfo(id int) map[string]any {
	var userInfo = make(map[string]any)
	ctx := context.Background()
	result, err := db.Rdb.Get(ctx, fmt.Sprintf("userInfo:%v", id)).Result()
	if err != nil {
		return userInfo
	}
	_ = json.Unmarshal([]byte(result), &userInfo)
	return userInfo
}

// SetUserInfo 设置用户信息
func setRedisUserInfo(user models.User) {
	ctx := context.Background()
	userData := map[string]any{
		"id":           db.GetStrID(user.ID),
		"username":     user.Username,
		"nickname":     user.Nickname,
		"avatar":       user.Avatar,
		"background":   user.Background,
		"introduction": user.Introduction,
		"coins":        user.Coins,
		"exp":          user.Exp,
		"level":        user.Level,
		"createdAt":    user.CreatedAt,
	}
	jsonData, _ := json.Marshal(userData)
	db.Rdb.Set(ctx, fmt.Sprintf("userInfo:%v", user.ID), jsonData, time.Hour*1)
}
