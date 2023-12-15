package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"log"
	"os"
	"time"
)

// GetUserInfo 获取用户信息
func GetUserInfo(userId uint) any {
	ctx := context.Background()
	userData, err := db.Rdb.Get(ctx, fmt.Sprintf("user_info:%v", userId)).Result()
	if err == nil {
		var data map[string]interface{}
		_ = json.Unmarshal([]byte(userData), &data)
		return data
	} else {
		var user models.User
		db.DB.Where("id = ?", userId).Find(&user)
		var data = map[string]interface{}{
			"id":       db.GetStrID(user.ID),
			"avatar":   user.Avatar,
			"username": user.Username,
			"nickname": user.Nickname,
		}
		db.Rdb.Set(ctx, fmt.Sprintf("user_info:%v", user.ID), data, time.Hour*2)

		return data
	}
}

func GetUserList(userIds []uint) []models.User {
	var users []models.User
	if err := db.DB.Model(models.User{}).Where("id IN ?", userIds).Find(&users).Error; err != nil {
		return make([]models.User, 0)
	}
	return users
}

// 确保所需目录存在
func EnsureDir(dirName string) {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
}
