package group

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
)

var name = "group_list"

// CreateGroup 创建群组
func CreateGroup(groupId uint, name string) error {
	ctx := context.Background()
	data := map[string]any{
		"id":   groupId,
		"name": name,
	}
	jsonData, _ := json.Marshal(data)
	_, err := db.Rdb.HSet(ctx, models.GroupListKey, groupId, jsonData).Result()
	if err != nil {
		return err
	}
	return nil
}

// GetGroupKey 获取群组的key
func GetGroupKey(groupId any) map[string]any {
	ctx := context.Background()
	key, _ := db.Rdb.HGet(ctx, models.GroupListKey, fmt.Sprintf("%v", groupId)).Result()
	var data map[string]any
	_ = json.Unmarshal([]byte(key), &data)
	return data
}
