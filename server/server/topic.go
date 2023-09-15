package server

import (
	"errors"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils"
	"log"
)

type TopicServer struct {
}
type Topic struct {
	ID      uint          `json:"id" `
	UserID  uint          `json:"userId"` // 用户ID
	Title   string        `json:"title"`  // 标题
	Tags    *db.IntArray  `json:"tags"`   // 标签
	Covers  *db.JSONSlice `json:"covers"` // 封面
	Type    int           `json:"type"`
	Summary string        `json:"summary"` // 摘要
	models.Model
	TopicDetail any `json:"topicDetail" gorm:"-"` // 详情
}

// Create 创建帖子
func (s *TopicServer) Create(post *models.Topic) (string, error) {
	if err := db.DB.Create(&post).Error; err != nil {
		return "", errors.New("unknown")
	}
	return utils.EncryptID(int(post.ID)), nil
}

// GetPostList 获取帖子列表
func (s *TopicServer) GetPostList(paging utils.Paging) (any, error) {
	var docs []Topic
	err := paging.SetPaging(db.DB).
		Order("id desc").
		Find(&docs).Error
	if err != nil {
		return nil, errors.New("unknown")
	}
	userIDs := make([]uint, 0)
	for _, doc := range docs {
		userIDs = append(userIDs, doc.UserID)
	}
	// 查询用户信息
	var users []models.User
	if err := db.DB.Model(models.User{}).Where("id IN ?", userIDs).Find(&users).Error; err != nil {
		return nil, errors.New("unknown")
	}
	userDetails := make(map[uint]map[string]any)
	for _, user := range users {
		userDetails[user.ID] = map[string]any{
			"id":       db.GetID(user.ID),
			"avatar":   user.Avatar,
			"username": user.Username,
			"nickname": user.Nickname,
		}
	}
	var topicList []map[string]any
	// 将详情填充到Post列表中
	for _, topic := range docs {
		topicList = append(topicList, map[string]any{
			"id":        db.GetID(topic.ID),
			"user":      userDetails[topic.UserID],
			"userId":    db.GetID(topic.UserID),
			"title":     topic.Title,
			"tags":      topic.Tags,
			"covers":    topic.Covers,
			"type":      topic.Type,
			"summary":   topic.Summary,
			"createdAt": topic.Model.CreatedAt,
		})
	}
	return topicList, nil
}

// GetTopicDetail 获取帖子详情
func (s *TopicServer) GetTopicDetail(id int) (any, error) {
	log.Println(id)
	var doc Topic
	err := db.DB.Where("id = ?", id).
		First(&doc).Error
	if err != nil {
		return nil, errors.New("unknown")
	}
	var detail map[string]any
	// 根据Type字段将Post IDs进行分组 1:默认 2:视频 3:图片 4:文本
	err = db.DB.Table(models.GetTopicType(doc.Type)).Where("topic_id = ?", id).
		Scan(&detail).Error
	if err != nil {
		return nil, errors.New("unknown")
	}
	// 查询用户信息
	user, err := models.GetUser(doc.UserID)
	if err != nil {
		return nil, errors.New("user_not_found")
	}
	return map[string]any{
		"id": db.GetID(doc.ID),
		"user": map[string]any{
			"id":       db.GetID(user.ID),
			"avatar":   user.Avatar,
			"username": user.Username,
			"nickname": user.Nickname,
		},
		"title":       doc.Title,
		"tags":        doc.Tags,
		"covers":      doc.Covers,
		"type":        doc.Type,
		"summary":     doc.Summary,
		"createdAt":   doc.Model.CreatedAt,
		"topicDetail": detail,
	}, nil
}
