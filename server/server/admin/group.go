package admin

import (
	"errors"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/klog"
)

type GroupServer struct {
}

// GroupInfo 获取用户组信息
func (g *GroupServer) GroupInfo(paging utils.Paging, name string) (any, error) {
	var doc []models.Group
	var total int64
	tx := db.DB
	if name != "" {
		tx = tx.Where("name like ?", "%"+name+"%")
	}
	tx.Model(&models.Group{}).Count(&total)
	err := paging.SetPaging(tx).Find(&doc).Error
	if err != nil {
		klog.Error("GroupInfo", err)
		return nil, errors.New("unknown")
	}
	return map[string]any{
		"list":  doc,
		"total": total,
	}, nil
}

// GroupAdd 添加用户组
func (g *GroupServer) GroupAdd(doc models.Group) (any, error) {
	err := db.DB.Create(&doc).Error
	if err != nil {
		klog.Error("GroupAdd", err)
		return nil, errors.New("json_error")
	}
	delRedisCache()
	return doc, nil
}

// GroupUpdate 修改用户组
func (g *GroupServer) GroupUpdate(doc models.Group) (any, error) {

	err := db.DB.Model(models.Group{}).Where("id = ?", doc.ID).Updates(&doc).Error
	if err != nil {
		klog.Error("GroupUpdate", err)
		return nil, errors.New("json_error")
	}
	delRedisCache()
	return doc, nil
}

// GroupDelete 删除用户组
func (g *GroupServer) GroupDelete(id string) (any, error) {
	err := db.DB.Delete(&models.Group{}, id).Error
	if err != nil {
		klog.Error("GroupDelete", err)
		return nil, errors.New("json_error")
	}
	delRedisCache()
	return nil, nil
}

// delRedisCache 删除redis缓存
func delRedisCache() {
	_ = db.Del(models.GroupListKey)
}
