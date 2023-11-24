package group

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
)

// CreateGroup 创建群组
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

	if len(data) == 0 {
		data = append(data, map[string]any{
			"id":    0,
			"order": 0,
			"name":  "默认",
		})
		var doc []models.Group
		db.DB.Find(&doc)
		for _, item := range doc {
			data = append(data, map[string]any{
				"id":    db.GetStrID(item.ID),
				"order": item.Order,
				"name":  item.Name,
			})
		}
		for _, item := range data {
			jsonData, _ := json.Marshal(item)
			_, _ = db.Rdb.HSet(ctx, "group", item["id"], jsonData).Result()

		}
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
