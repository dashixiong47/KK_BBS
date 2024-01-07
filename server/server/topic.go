package server

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/dashixiong47/KK_BBS/config"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/services"
	"github.com/dashixiong47/KK_BBS/services/group"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"github.com/dashixiong47/KK_BBS/utils/message"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
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
		for i, _ := range *attachments {
			attachment := &(*attachments)[i] // 正确解引用并获取切片中的元素
			attachment.TopicID = post.ID

		}
		if err := tx.Create(&attachments).Error; err != nil {
			tx.Rollback()
			klog.Error("Create Attachment", err)
			return "", errors.New("unknown")
		}
	}

	// 每次创建新帖子的时候，都会在Redis里设置该帖子的点赞默认为0
	err := services.CreateTopicLike(int(post.ID))
	if err != nil {
		tx.Rollback()
		return "", errors.New("unknown")
	}
	//
	err = services.AddTopicIntegral(post.UserID, post.ID)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	if err := tx.Commit().Error; err != nil {
		return "", errors.New("unknown")
	}
	go func() {
		var search models.Search
		marshal, _ := json.Marshal(post.Covers)
		search.Title = post.Title
		search.ID = db.GetStrID(post.ID)
		search.Type = strconv.Itoa(post.Type)
		search.Cover = string(marshal)
		search.GroupID = db.GetStrID(post.GroupID)
		if post.Type == 1 {
			search.Content = post.TopicBasic.Content
		}
		IndexPost(search.ID, search)
	}()
	return utils.EncryptID(int(post.ID)), nil
}

// GetTopicList 获取帖子列表
func (s *TopicServer) GetTopicList(_type, groupId string, userId, selfUserId uint, paging utils.Paging) (any, error) {
	var docs []Topic
	var count int64
	tx := db.DB
	tx = tx.
		Order("id desc").Model(&models.Topic{})
	if userId != 0 {
		tx = tx.Where("user_id = ?", userId)
	}
	if groupId != "-1" {
		tx = tx.Where("group_id = ?", groupId)
	}
	if _type != "-1" {
		tx = tx.Where("type = ?", _type)
	}
	tx.Count(&count)
	tx = paging.SetPaging(tx)
	err := tx.Find(&docs).Error
	if err != nil {
		klog.Error("GetTopicList", err)
		return nil, errors.New("unknown")
	}
	userIDs := make([]uint, 0)
	for _, doc := range docs {
		userIDs = append(userIDs, doc.UserID)
	}
	followState := make(map[uint]bool)
	if selfUserId != 0 {
		// 关注状态
		followState = services.GetFollowStatus(selfUserId, userIDs)
	}
	// 查询用户信息
	userDetails := make(map[uint]any)
	for _, user := range services.GetUserList(userIDs) {
		userDetails[user.ID], err = services.GetUserDetailInfo(int(user.ID))
	}
	var topicList []map[string]any
	// 将详情填充到Post列表中
	for _, topic := range docs {
		if selfUserId != 0 {
			userDetails[topic.UserID].(map[string]any)["isFollow"] = followState[topic.UserID]
		}
		topicList = append(topicList, map[string]any{
			"id":           db.GetStrID(topic.ID),
			"user":         userDetails[topic.UserID],
			"userId":       db.GetStrID(topic.UserID),
			"title":        topic.Title,
			"tags":         topic.Tags,
			"covers":       topic.Covers,
			"type":         topic.Type,
			"summary":      topic.Summary,
			"createdAt":    topic.Model.CreatedAt,
			"view":         services.GetTopicViewCount(topic.ID),
			"comment":      services.GetTopicCommentCount(topic.ID),
			"commentState": services.IsTopicComment(topic.ID, selfUserId),
			"like":         services.GetTopicLikeCount(topic.ID),
			"likeState":    services.IsTopicLike(topic.ID, selfUserId),
			"group":        group.GetGroupKey(topic.GroupID),
			"collect":      services.GetTopicCollectCount(topic.UserID),
			"collectState": services.IsTopicCollect(topic.ID, selfUserId),
		})
	}
	return map[string]any{
		"list":  topicList,
		"total": count,
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
	// 根据Type字段将Post IDs进行分组 1:默认 2:图片 3:视频 4:文本
	err = db.DB.Table(models.GetTopicType(doc.Type)).Where("topic_id = ?", topicId).
		Scan(&detail).Error
	if err != nil {
		return nil, errors.New("unknown")
	}
	if doc.Type == 1 && detail["hidden"] == int16(1) {
		var count int64
		db.DB.Model(models.Comment{}).Where("topic_id = ?", topicId).
			Where("user_id = ?", userId).Count(&count)
		if count == 0 {
			detail["hidden_content"] = "该内容已隐藏，需要评论后才能查看"

		}
	} else if doc.Type == 2 {
		images := make([]map[string]any, 0)
		err := json.Unmarshal([]byte(detail["images"].(string)), &images)
		if err != nil {
			log.Println("json.Unmarshal", err)
		}
		detail["images"] = images

	} else if doc.Type == 3 {
		log.Println(detail)
		videos := make([]map[string]any, 0)
		err := json.Unmarshal([]byte(detail["videos"].(string)), &videos)
		if err != nil {
			log.Println("json.Unmarshal", err)
		}
		detail["videos"] = videos

	} else if doc.Type == 4 {
		texts := make([]map[string]any, 0)
		err := json.Unmarshal([]byte(detail["texts"].(string)), &texts)
		if err != nil {
			log.Println("json.Unmarshal", err)
		}
		detail["texts"] = texts

	}
	err = services.TopicViewPlus(uint(topicId), doc.UserID)
	if err != nil {
		klog.Error("TopicViewPlus", err)
	}
	// 用户信息
	var userInfo, _ = services.GetUserDetailInfo(int(doc.UserID))
	return map[string]any{
		"id":           db.GetStrID(doc.ID),
		"user":         userInfo,
		"title":        doc.Title,
		"tags":         doc.Tags,
		"covers":       doc.Covers,
		"type":         doc.Type,
		"view":         services.GetTopicViewCount(uint(topicId)),
		"like":         services.GetTopicLikeCount(uint(topicId)),
		"likeState":    services.IsTopicLike(uint(topicId), doc.UserID),
		"collect":      services.GetTopicCollectCount(doc.UserID),
		"collectState": services.IsTopicCollect(uint(topicId), uint(userId)),
		"comment":      services.GetTopicCommentCount(uint(topicId)),
		"commentState": services.IsTopicComment(uint(topicId), uint(userId)),
		"summary":      doc.Summary,
		"createdAt":    doc.Model.CreatedAt,
		"topicDetail":  detail,
	}, nil
}

// LikeTopic 点赞帖子
func (s *TopicServer) LikeTopic(topicId, userId int) error {
	// 检查帖子是否存在
	if !services.IsTopic(topicId) {
		return errors.New("topic_not_found")
	}

	var topicLike models.TopicLike
	err := db.DB.Unscoped().Where("topic_id = ? AND user_id = ?", topicId, userId).First(&topicLike).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("database_query_error")
	}

	tx := db.DB.Begin()

	// 更新或创建点赞记录
	if topicLike.ID != 0 {
		// 判断是取消点赞还是标记为已点赞
		if topicLike.DeletedAt.Valid {
			// 已经被软删除，需要恢复（取消软删除）
			if err := tx.Unscoped().Model(&topicLike).Update("deleted_at", gorm.DeletedAt{}).Error; err != nil {
				tx.Rollback()
				return errors.New("restore_like_error")
			}
		} else {
			// 标记为软删除（即取消点赞）
			if err := tx.Model(&topicLike).Update("deleted_at", time.Now()).Error; err != nil {
				tx.Rollback()
				return errors.New("delete_like_error")
			}
		}
	} else {
		// 创建新的点赞记录
		topicLike = models.TopicLike{TopicID: uint(topicId), UserID: uint(userId)}
		if err := tx.Create(&topicLike).Error; err != nil {
			tx.Rollback()
			return errors.New("create_like_error")
		}

		// 增加帖子点赞计数
		if err := services.TopicLickPlus(topicId); err != nil {
			tx.Rollback()
			return errors.New("like_count_error")
		}
		err := sendLikeMessageIfNecessary(topicId, userId)
		klog.Error("LikeTopic", err)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return errors.New("transaction_commit_error")
	}

	return nil
}

// sendLikeMessageIfNecessary 发送点赞消息（如果需要）
func sendLikeMessageIfNecessary(topicId, userId int) error {
	var topicUserId uint
	err := db.DB.Model(models.Topic{}).Select("user_id").Where("id = ?", topicId).
		Row().Scan(&topicUserId)
	if err != nil {
		klog.Error("LikeTopic", err)
		return nil // 不影响主要功能
	}
	if userId != int(topicUserId) {
		err = message.SendLikeMessage(uint(userId), topicUserId, uint(topicId), "")
		if err != nil {
			klog.Error("LikeTopic", err)
		}
	}
	return nil
}

// CollectTopic 收藏帖子
func (s *TopicServer) CollectTopic(topicId, userId int) error {
	// 检查帖子是否存在
	if !services.IsTopic(topicId) {
		return errors.New("topic_not_found")
	}

	var topicCollect models.Collection
	err := db.DB.Unscoped().Where("topic_id = ? AND user_id = ?", topicId, userId).First(&topicCollect).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("database_query_error")
	}

	tx := db.DB.Begin()

	// 更新或创建收藏记录
	if topicCollect.ID != 0 {
		// 判断是取消收藏还是标记为已收藏
		if topicCollect.DeletedAt.Valid {
			// 已经被软删除，需要恢复（取消软删除）
			if err := tx.Unscoped().Model(&topicCollect).Update("deleted_at", gorm.DeletedAt{}).Error; err != nil {
				tx.Rollback()
				return errors.New("restore_like_error")
			}
		} else {
			// 标记为软删除（即取消收藏）
			if err := tx.Model(&topicCollect).Update("deleted_at", time.Now()).Error; err != nil {
				tx.Rollback()
				return errors.New("delete_like_error")
			}
		}
	} else {
		// 创建新的收藏记录
		topicCollect = models.Collection{TopicID: uint(topicId), UserID: uint(userId)}
		if err := tx.Create(&topicCollect).Error; err != nil {
			tx.Rollback()
			return errors.New("create_collect_error")
		}
		err := sendCollectMessageIfNecessary(topicId, userId, topicCollect)
		klog.Error("CollectTopic", err)
	}

	// 清除相关缓存
	services.ClearTopicCollectCountCache(uint(userId))

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return errors.New("transaction_commit_error")
	}

	// 发送收藏消息（如果有必要）
	return nil
}

// sendCollectMessageIfNecessary 发送收藏消息（如果需要）
func sendCollectMessageIfNecessary(topicId, userId int, collect models.Collection) error {
	var topicUserId uint
	err := db.DB.Model(models.Topic{}).Select("user_id").Where("id = ?", topicId).
		Row().Scan(&topicUserId)
	if err != nil {
		klog.Error("CollectTopic", err)
		return nil // 不影响主要功能
	}

	if userId != int(topicUserId) {
		err = message.SendCollectMessage(uint(userId), topicUserId, uint(topicId), "")
		if err != nil {
			klog.Error("CollectTopic", err)
		}
	}
	return nil
}

// GetCollectList 获取收藏列表
func (s *TopicServer) GetCollectList(userId int, paging utils.Paging) (any, error) {
	var docs []models.Topic
	var count int64
	tx := db.DB.Model(&models.Topic{}).
		Joins("left join kk_collection on kk_collection.topic_id = kk_topic.id").
		Where("kk_collection.user_id = ?", userId).
		Where("kk_collection.deleted_at is null")
	tx.Count(&count)
	tx = paging.SetPaging(tx)
	err := tx.Find(&docs).Error
	if err != nil {
		klog.Error("GetCollectList", err)
		return nil, err
	}
	userIDs := make([]uint, 0)
	for _, doc := range docs {
		userIDs = append(userIDs, doc.UserID)
	}
	// 查询用户信息
	userDetails := make(map[uint]map[string]any)
	for _, user := range services.GetUserList(userIDs) {
		userDetails[user.ID] = map[string]any{
			"id":       db.GetStrID(user.ID),
			"avatar":   user.Avatar,
			"username": user.Username,
			"nickname": user.Nickname,
		}
	}
	var topicList []map[string]any
	// 将详情填充到Post列表中
	for _, topic := range docs {
		topicList = append(topicList, map[string]any{
			"id":           db.GetStrID(topic.ID),
			"user":         userDetails[topic.UserID],
			"userId":       db.GetStrID(topic.UserID),
			"title":        topic.Title,
			"tags":         topic.Tags,
			"covers":       topic.Covers,
			"type":         topic.Type,
			"summary":      topic.Summary,
			"createdAt":    topic.Model.CreatedAt,
			"view":         services.GetTopicViewCount(topic.ID),
			"comment":      services.GetTopicCommentCount(topic.ID),
			"commentState": services.IsTopicComment(topic.ID, uint(userId)),
			"like":         services.GetTopicLikeCount(topic.ID),
			"likeState":    services.IsTopicLike(topic.ID, uint(userId)),
			"group":        group.GetGroupKey(topic.GroupID),
			"collect":      services.GetTopicCollectCount(topic.UserID),
			"collectState": services.IsTopicCollect(topic.ID, uint(userId)),
		})
	}
	return map[string]any{
		"list":  topicList,
		"total": count,
	}, nil
}

// IndexPost 插入es
func IndexPost(id string, post models.Search) {
	_, err := db.EDB.Index().
		Index(config.SettingsConfig.Es.Index).
		Id(id).
		BodyJson(post).
		Do(context.Background())
	if err != nil {
		klog.Error("IndexPost", err)
	}
}
