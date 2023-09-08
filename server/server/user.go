package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"time"
)

type UserServer struct {
}

func (u *UserServer) GetUserInfo(id int) (models.User, error) {
	var userInfo models.User
	// 先从缓存中获取
	info := getRedisUserInfo(id)
	if info.ID != 0 {
		return info, nil
	}
	// 如果缓存中没有，再从数据库中获取
	if err := db.DB.Where("id = ?", id).First(&userInfo).Error; err != nil {
		return userInfo, errors.New("user_not_found")
	}
	// 将数据存到缓存中
	setRedisUserInfo(userInfo)
	return userInfo, nil
}

func getRedisUserInfo(id int) models.User {
	var userInfo models.User
	ctx := context.Background()
	result, err := db.Rdb.Get(ctx, fmt.Sprintf("userInfo:%v", id)).Result()
	if err != nil {
		return userInfo
	}
	_ = json.Unmarshal([]byte(result), &userInfo)
	return userInfo
}
func setRedisUserInfo(user models.User) {
	ctx := context.Background()
	data, _ := json.Marshal(user)
	db.Rdb.Set(ctx, fmt.Sprintf("userInfo:%v", user.ID), data, time.Hour*1)

}
