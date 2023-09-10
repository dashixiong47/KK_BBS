package server

import (
	"errors"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils"
)

type PostServer struct {
}
type Post struct {
	ID     uint          `json:"id" `
	UserID uint          `json:"userId"` // 用户ID
	Title  string        `json:"title"`  // 标题
	Tags   *db.IntArray  `json:"tags"`   // 标签
	Covers *db.JSONSlice `json:"covers"` // 封面
	Type   int           `json:"type"`
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
	//for i, post := range posts {
	//	posts[i].ID = db.GetID(post.ID)
	//}
	return posts, nil
}
