package server

import (
	"errors"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils"
)

type PostServer struct {
}

// Create 创建帖子
func (s *PostServer) Create(post *models.Post) error {
	if err := db.DB.Create(&post).Error; err != nil {
		return errors.New("unknown_error")
	}
	return nil
}

// GetPostList 获取帖子列表
func (s *PostServer) GetPostList(paging utils.Paging) ([]models.Post, error) {
	var posts []models.Post
	err := paging.SetPaging(db.DB).Find(&posts).Error
	if err != nil {
		return nil, errors.New("unknown_error")
	}
	return posts, nil
}
