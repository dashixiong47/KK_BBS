package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
	"strconv"
)

func UserIntegral(userId uint) int {
	ctx := context.Background()
	result, err := db.Rdb.HGet(ctx, "userIntegral", fmt.Sprintf("%v", userId)).Result()
	if err != nil {
		return 0
	}
	integral, err := strconv.Atoi(result)
	if err != nil {
		return 0
	}
	return integral
}

func RemoveUserRedis(userId int) error {
	err := db.Rdb.Del(context.Background(), fmt.Sprintf("userInfo:%v", userId)).Err()
	if err != nil {
		return errors.New("update_user_info_error")
	}
	return nil
}
