package group

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
)

// CreateGroup 创建帖子点赞
func CreateGroup(groupId uint, name string) error {
	ctx := context.Background()
	data := map[string]any{
		"id":   groupId,
		"name": name,
	}
	jsonData, _ := json.Marshal(data)
	_, err := db.Rdb.HSet(ctx, "group", groupId, jsonData).Result()
	if err != nil {
		return err
	}
	return nil
}

// GetAllGroup 获取所有的群组
func GetAllGroup() []map[string]any {
	ctx := context.Background()
	groups, _ := db.Rdb.HGetAll(ctx, "group").Result()
	var data = make([]map[string]any, 0)
	for _, group := range groups {
		var item map[string]any
		_ = json.Unmarshal([]byte(group), &item)
		data = append(data, item)
	}
	return data
}

// GetGroupKey 获取群组的key
func GetGroupKey(groupId any) map[string]any {
	ctx := context.Background()
	key, _ := db.Rdb.HGet(ctx, "group", fmt.Sprintf("%v", groupId)).Result()
	var data map[string]any
	_ = json.Unmarshal([]byte(key), &data)
	return data
}
