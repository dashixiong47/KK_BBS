package server

import (
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
)

type UserServer struct {
}

func (u *UserServer) UserInfo(id int) (models.User, error) {
	var userInfo models.User
	db.DB.Where("id = ?", id).First(&userInfo)
	return userInfo, nil
}
