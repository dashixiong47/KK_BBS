package server

import (
	"errors"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/services"
)

type UserServer struct {
}

// GetUserInfo 获取用户信息
func (u *UserServer) GetUserInfo(id int) (any, error) {
	return services.GetUserDetailInfo(id)
}

// UpdateUserInfo 修改用户信息
func (u *UserServer) UpdateUserInfo(id int, nickname, avatar, background, introduction string) (any, error) {

	tx := db.DB.Begin().Model(models.User{}).Where("id = ?", id)
	if nickname != "" {
		tx = tx.Update("nickname", nickname)
	}
	if avatar != "" {
		tx = tx.Update("avatar", avatar)
	}
	if background != "" {
		tx = tx.Update("background", background)
	}
	if introduction != "" {
		tx = tx.Update("introduction", introduction)
	}
	err := services.RemoveUserRedis(id)
	if err != nil {
		return nil, err
	}
	err = tx.Commit().Error
	if err != nil {
		return nil, errors.New("update_user_info_error")
	}
	return nil, nil
}
