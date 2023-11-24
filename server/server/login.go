package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/jwt"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"strings"
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
		"id":       db.GetStrID(user.ID),
		"avatar":   user.Avatar,
		"username": user.Username,
		"nickname": user.Nickname,
	}, nil
}

// Register 注册
func (s *LoginServer) Register(email, vCode string) (any, error) {
	//查询
	var user models.User
	ctx := context.Background()
	key := fmt.Sprintf("verifCode:%s", email)
	result, err := db.Rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, errors.New("email_not_found")
	}
	if result != vCode {
		return nil, errors.New("vcode_error")
	}
	// 查询是否已经注册
	err = db.DB.Model(user).Where("email = ?", email).First(&user).Error
	if err != nil {
		pwd := utils.GenerateRandomString(6)
		user.Email = email
		user.Username = email
		user.Nickname = strings.Split(email, "@")[0]
		user.Status = 1
		user.Password = utils.MD5(pwd)
		err = db.DB.Create(&user).Error
		if err != nil {
			return nil, errors.New("register_error")
		}
	}
	_, err = db.Rdb.Del(ctx, key).Result()
	if err != nil {
		klog.Error("Del VerificationCode", err)
	}
	token, err := jwt.CreateToken(user)
	if err != nil {
		return nil, errors.New("token_error")
	}
	return map[string]string{
		"token":    token,
		"id":       db.GetStrID(user.ID),
		"avatar":   user.Avatar,
		"username": user.Username,
		"nickname": user.Nickname,
	}, nil
}
