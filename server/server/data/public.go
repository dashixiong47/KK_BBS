package data

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
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
			"id":       db.GetID(user.ID),
			"avatar":   user.Avatar,
			"username": user.Username,
			"nickname": user.Nickname,
		}
		db.Rdb.Set(ctx, fmt.Sprintf("user_info:%v", user.ID), data, time.Hour*2)

		return data
	}
}
