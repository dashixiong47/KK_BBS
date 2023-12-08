package server

import (
	"errors"
	"github.com/dashixiong47/KK_BBS/data"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"gorm.io/gorm"
	"time"
)

type FollowServer struct {
}

func (f *FollowServer) CreateFollow(userId, followId int) error {
	// 检查用户是否存在
	if !data.IsUser(userId) {
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
