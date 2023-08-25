package server

import (
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils/jwt"
)

type LoginServer struct {
}

func (s *LoginServer) Login(username, password string) (any, error) {
	var user models.User
	err := db.DB.Model(user).
		Where("username = ? AND password = ?", username, password).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	token, err := jwt.CreateToken(user)
	if err != nil {
		return nil, err
	}

	return token, nil
}
