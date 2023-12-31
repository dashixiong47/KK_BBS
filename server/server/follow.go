package server

import (
	"errors"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/services"
	"github.com/dashixiong47/KK_BBS/utils"
	"gorm.io/gorm"
	"log"
	"time"
)

type FollowServer struct {
}

func (f *FollowServer) CreateFollow(userId, followId int) error {

	if userId == followId {
		return errors.New("can_not_follow_self")
	}
	// 检查用户是否存在
	if !services.IsUser(userId) {
		return errors.New("user_not_found")
	}

	var follow models.Follow
	err := db.DB.Unscoped().Where("user_id = ? AND follow_id = ?", userId, followId).First(&follow).Error
	// 处理查询错误
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("database_query_error")
	}

	tx := db.DB.Begin()

	// 更新或创建关注记录
	if follow.ID != 0 {
		// 更新关注记录
		if follow.DeletedAt.Valid {
			// 已经被软删除，需要恢复（取消软删除）
			if err := tx.Unscoped().Model(&follow).Update("deleted_at", gorm.DeletedAt{}).Error; err != nil {
				tx.Rollback()
				return errors.New("restore_like_error")
			}
		} else {
			// 未被软删除，需要软删除
			if err := tx.Model(&follow).Update("deleted_at", time.Now()).Error; err != nil {
				tx.Rollback()
				return errors.New("delete_like_error")
			}
		}
	} else {
		// 创建新的关注记录
		follow = models.Follow{UserID: uint(userId), FollowID: uint(followId)}
		if err := tx.Create(&follow).Error; err != nil {
			tx.Rollback()
			return errors.New("create_follow_error")
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return errors.New("transaction_commit_error")
	}
	return nil
}

// GetFollowList 获取关注列表
func (f *FollowServer) GetFollowList(userId int, paging utils.Paging) (any, error) {
	if !services.IsUser(userId) {
		return nil, errors.New("user_not_found")
	}
	var followIdList []uint
	var count int64
	if err := paging.SetPaging(db.DB).Model(models.Follow{}).
		Select("follow_id").
		Where("user_id = ?", userId).Scan(&followIdList).
		Count(&count).Error; err != nil {
		return nil, errors.New("database_query_error")
	}
	followStatus := services.GetFollowStatus(uint(userId), followIdList)
	log.Println(followIdList)
	var userList = make([]any, 0)
	for _, followId := range followIdList {

		info, _ := services.GetUserDetailInfo(int(followId))
		info.(map[string]any)["isFollow"] = followStatus[followId]
		userList = append(userList, info)
	}
	var data = make(map[string]any)
	data["list"] = userList
	data["total"] = count
	return data, nil
}

// GetFansList 获取粉丝列表
func (f *FollowServer) GetFansList(userId int, paging utils.Paging) (any, error) {
	if !services.IsUser(userId) {
		return nil, errors.New("user_not_found")
	}
	var fansIdList []uint
	var count int64
	if err := paging.SetPaging(db.DB).
		Model(models.Follow{}).
		Select("user_id").
		Where("follow_id = ?", userId).
		Scan(&fansIdList).
		Count(&count).Error; err != nil {
		return nil, errors.New("database_query_error")
	}
	followStatus := services.GetFollowStatus(uint(userId), fansIdList)
	var userList = make([]any, 0)
	for _, fansId := range fansIdList {
		info, _ := services.GetUserDetailInfo(int(fansId))
		info.(map[string]any)["isFollow"] = followStatus[fansId]
		userList = append(userList, info)
	}
	var data = make(map[string]any)
	data["list"] = userList
	data["total"] = count
	return data, nil
}
