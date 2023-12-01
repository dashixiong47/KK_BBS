package server

import (
	"context"
	"encoding/json"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
)

type GroupServer struct{}

func (g *GroupServer) GroupInfo() (any, error) {
	ctx := context.Background()
	groups, _ := db.Rdb.HGetAll(ctx, models.GroupListKey).Result()
	var data = make([]map[string]any, 0)
	for _, _group := range groups {
		var item map[string]any
		_ = json.Unmarshal([]byte(_group), &item)
		data = append(data, item)
	}

	if len(data) == 0 {
		var doc []models.Group
		db.DB.Find(&doc)
		for _, item := range doc {
			data = append(data, map[string]any{
				"id":    item.ID,
				"order": item.Order,
				"icon":  item.Icon,
				"name":  item.Name,
			})
		}
		for _, item := range data {
			jsonData, _ := json.Marshal(item)
			_, _ = db.Rdb.HSet(ctx, models.GroupListKey, item["id"], jsonData).Result()

		}
	}

	return data, nil
	//return group.GetAllGroup(), nil
}
