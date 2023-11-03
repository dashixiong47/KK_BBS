package server

import (
	"errors"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/server/data"
	"github.com/dashixiong47/KK_BBS/server/data/group"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"log"
)

type TopicServer struct {
}
type Topic struct {
	ID      uint          `json:"id" `
	UserID  uint          `json:"userId"`  // 用户ID
	GroupID uint          `json:"groupId"` // 分组
	Title   string        `json:"title"`   // 标题
	Tags    *db.IntArray  `json:"tags"`    // 标签
	Covers  *db.JSONSlice `json:"covers"`  // 封面
	Type    int           `json:"type"`
	Summary string        `json:"summary"` // 摘要
	models.Model
	TopicDetail any `json:"topicDetail" gorm:"-"` // 详情
}

// Create 创建帖子
func (s *TopicServer) Create(post *models.Topic, attachments *[]models.Attachment) (string, error) {
	tx := db.DB.Begin()
	if err := tx.Create(&post).Error; err != nil {
		tx.Rollback()
		return "", errors.New("unknown")
	}

	if len(*attachments) > 0 {
		for _, attachment := range *attachments {
			attachment.TopicID = post.ID
		}
		if err := tx.Debug().Create(&attachments).Error; err != nil {
			tx.Rollback()
			klog.Error("Create Attachment", err)
			return "", errors.New("unknown")
		}
	}
	// 每次创建新帖子的时候，都会在Redis里设置该帖子的点赞默认为0
	err := data.CreateTopicLike(int(post.ID))
	if err != nil {
		tx.Rollback()
		return "", errors.New("unknown")
	}
	if err := tx.Commit().Error; err != nil {
		return "", errors.New("unknown")
	}
	return utils.EncryptID(int(post.ID)), nil
}

// GetTopicList 获取帖子列表
func (s *TopicServer) GetTopicList(_type string, userId uint, paging utils.Paging) (any, error) {
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
	userDetails := make(map[uint]map[string]any)
	for _, user := range data.GetUserList(userIDs) {
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
		log.Println(topic.GroupID)
		topicList = append(topicList, map[string]any{
			"id":           db.GetID(topic.ID),
			"user":         userDetails[topic.UserID],
			"userId":       db.GetID(topic.UserID),
			"title":        topic.Title,
			"tags":         topic.Tags,
			"covers":       topic.Covers,
			"type":         topic.Type,
			"summary":      topic.Summary,
			"createdAt":    topic.Model.CreatedAt,
			"view":         data.GetTopicViewCount(topic.ID),
			"comment":      data.GetTopicCommentCount(topic.ID),
			"commentState": data.IsTopicComment(topic.ID, userId),
			"like":         data.GetTopicLikeCount(topic.ID),
			"likeState":    data.IsTopicLike(topic.ID, userId),
			"group":        group.GetGroupKey(topic.GroupID),
			"collect":      data.GetTopicCollectCount(topic.UserID),
			"collectState": data.IsTopicCollect(topic.ID, userId),
		})
	}
	return map[string]any{
		"list":  topicList,
		"total": data.GetTopicCount(),
	}, nil
}

// GetTopicDetail 获取帖子详情
func (s *TopicServer) GetTopicDetail(topicId, userId int) (any, error) {
	var doc Topic
	err := db.DB.Where("id = ?", topicId).
		First(&doc).Error
	if err != nil {
		return nil, errors.New("unknown")
	}
	var detail map[string]any
	// 根据Type字段将Post IDs进行分组 1:默认 2:视频 3:图片 4:文本
	err = db.DB.Table(models.GetTopicType(doc.Type)).Where("topic_id = ?", topicId).
		Scan(&detail).Error
	if err != nil {
		return nil, errors.New("unknown")
	}
	// 查询用户信息
	user, err := models.GetUser(doc.UserID)
	if err != nil {
		return nil, errors.New("user_not_found")
	}
	err = data.TopicViewPlus(uint(topicId), doc.UserID)
	if err != nil {
		klog.Error("TopicViewPlus", err)
	}

	return map[string]any{
		"id": db.GetID(doc.ID),
		"user": map[string]any{
			"id":       db.GetID(user.ID),
			"avatar":   user.Avatar,
			"username": user.Username,
			"nickname": user.Nickname,
		},
		"title":        doc.Title,
		"tags":         doc.Tags,
		"covers":       doc.Covers,
		"type":         doc.Type,
		"view":         data.GetTopicViewCount(uint(topicId)),
		"like":         data.GetTopicLikeCount(uint(topicId)),
		"likeState":    data.IsTopicLike(uint(topicId), doc.UserID),
		"collect":      data.GetTopicCollectCount(doc.UserID),
		"collectState": data.IsTopicCollect(uint(topicId), uint(userId)),
		"comment":      data.GetTopicCommentCount(uint(topicId)),
		"commentState": data.IsTopicComment(uint(topicId), uint(userId)),
		"summary":      doc.Summary,
		"createdAt":    doc.Model.CreatedAt,
		"topicDetail":  detail,
	}, nil
}

// LikeTopic 点赞
func (s *TopicServer) LikeTopic(topicId, userId int) error {
	var topicLike models.TopicLike
	state := data.IsTopic(topicId)
	if !state {
		return errors.New("is_not_topic")
	}
	// 查询是否已经点赞
	db.DB.Unscoped().Where("topic_id = ? and user_id = ?", topicId, userId).
		First(&topicLike)
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 如果已经点赞 则取消点赞
	if topicLike.ID != 0 && topicLike.DeletedAt.Valid {
		if err := tx.Model(&topicLike).
			Unscoped().
			Update("deleted_at", nil).Error; err != nil {
			tx.Rollback()
			return errors.New("unknown")
		}
		err := data.TopicLickPlus(topicId)
		if err != nil {
			tx.Rollback()
			return errors.New("unknown")
		}
		// 如果是正常的记录，则进行软删除
	} else if topicLike.ID != 0 && !topicLike.DeletedAt.Valid {
		if err := tx.Delete(&topicLike).Error; err != nil {
			tx.Rollback()
			return errors.New("unknown")
		}
		err := data.TopicLickMinus(topicId)
		if err != nil {
			tx.Rollback()
			return errors.New("unknown")
		}
		// 不存在该记录，创建新记录
	} else {
		// 未点赞
		topicLike = models.TopicLike{
			TopicID: uint(topicId),
			UserID:  uint(userId),
		}
		err := db.DB.Create(&topicLike).Error
		if err != nil {
			tx.Rollback()
			return errors.New("error")
		}
		err = data.TopicLickPlus(topicId)
		if err != nil {
			tx.Rollback()
			return errors.New("unknown")
		}
		return nil
	}

	if err := tx.Commit().Error; err != nil {
		return errors.New("unknown")
	}
	return nil
}

// CollectTopic 收藏帖子
func (s *TopicServer) CollectTopic(topicId, userId int) error {
	var topicCollect models.Collection
	state := data.IsTopic(topicId)
	if !state {
		return errors.New("is_not_topic")
	}
	// 查询是否已经收藏
	db.DB.Unscoped().Where("topic_id = ? and user_id = ?", topicId, userId).
		First(&topicCollect)
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// 如果已经收藏 则取消收藏
	if topicCollect.ID != 0 && topicCollect.DeletedAt.Valid {
		if err := tx.Model(&topicCollect).
			Unscoped().
			Update("deleted_at", nil).Error; err != nil {
			tx.Rollback()
			return errors.New("unknown")
		}
		// 如果是正常的记录，则进行软删除
	} else if topicCollect.ID != 0 && !topicCollect.DeletedAt.Valid {
		if err := tx.Delete(&topicCollect).Error; err != nil {
			tx.Rollback()
			return errors.New("unknown")
		}
		// 不存在该记录，创建新记录
	} else {
		// 未收藏
		topicCollect = models.Collection{
			TopicID: uint(topicId),
			UserID:  uint(userId),
		}
		err := db.DB.Create(&topicCollect).Error
		if err != nil {
			tx.Rollback()
			return errors.New("error")
		}
		return nil
	}
	data.ClearTopicCollectCountCache(uint(userId))
	if err := tx.Commit().Error; err != nil {
		return errors.New("unknown")
	}
	return nil
}
