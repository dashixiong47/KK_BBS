package server

import (
	"errors"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils/jwt"
)

type LoginServer struct {
}

// Login 登录
func (s *LoginServer) Login(username, password string) (any, error) {
	var user models.User
	err := db.DB.Model(user).
		Where("username = ? AND password = ?", username, password).
		First(&user).Error
	if err != nil {
		return nil, errors.New("username_or_password_error")
	}

	token, err := jwt.CreateToken(user)
	if err != nil {
		return nil, errors.New("token_error")
	}
	return map[string]string{
		"token":    token,
		"id":       db.GetID(user.ID),
		"avatar":   user.Avatar,
		"username": user.Username,
		"nickname": user.Nickname,
	}, nil
}
