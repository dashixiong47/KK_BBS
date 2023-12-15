package server

import (
	"errors"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/services"
	"github.com/dashixiong47/KK_BBS/utils/klog"
)

type AttachmentServer struct {
}

// List 获取附件列表
func (s *AttachmentServer) List(topId int, userId uint) ([]map[string]any, error) {
	// 定义一个结构体来存储查询结果
	type AttachmentInfo struct {
		models.Attachment
		Purchased bool `gorm:"column:purchased"`
	}

	var attachmentsInfo []AttachmentInfo

	// 执行原生SQL查询
	err := db.DB.Debug().
		Table("kk_attachment").
		Select("kk_attachment.*, CASE WHEN kk_integral_log.id IS NOT NULL THEN true ELSE false END as purchased").
		Joins("left join kk_integral_log on kk_attachment.id = kk_integral_log.source_id AND kk_integral_log.user_id = ? AND kk_integral_log.type = ? AND kk_integral_log.source = ?", userId, models.IntegralTypeReduce, models.SourceTypePurchaseAttachment).
		Where("kk_attachment.topic_id = ?", topId).
		Find(&attachmentsInfo).Error

	if err != nil {
		return nil, errors.New("attachment_not_found")
	}

	// 构建最终的文档列表
	docs := make([]map[string]any, 0, len(attachmentsInfo))
	for _, info := range attachmentsInfo {
		isSelf := info.UserID == userId
		doc := make(map[string]any)
		doc["id"] = db.GetStrID(info.ID)
		doc["topicId"] = db.GetStrID(info.TopicID)
		doc["userId"] = db.GetStrID(info.UserID)
		doc["name"] = info.Name
		doc["createdAt"] = info.CreatedAt
		doc["type"] = info.Type
		doc["order"] = info.Order
		doc["hidden"] = info.Hidden
		doc["coins"] = info.Coins
		doc["purchased"] = info.Purchased

		// 根据附件类型和用户是否有权访问来设置文件信息
		if info.Type == 1 || info.Coins == 0 || info.Purchased {
			doc["fileType"] = info.FileType
			doc["fileSize"] = info.FileSize
			// 如果用户已经购买了附件或者附件是免费的，则提供文件URL
			if info.Purchased || info.Coins == 0 || isSelf {
				doc["fileUrl"] = info.FileUrl
				doc["status"] = 1
			}
		}

		// 对于类型为2的附件，如果用户有权访问或者附件是免费的，提供网盘信息
		if info.Type == 2 && (info.Purchased || info.Coins == 0) || isSelf {
			doc["netDisk"] = info.NetDiskType
			doc["netDiskUrl"] = info.NetDiskUrl
			doc["netDiskCode"] = info.NetDiskCode
			doc["status"] = 1
		}

		// 添加备注信息
		doc["remake"] = info.Remake

		docs = append(docs, doc)
	}

	return docs, nil
}

// Buy 购买附件
func (s *AttachmentServer) Buy(attachmentId int, userId uint) error {

	var attachment models.Attachment
	err := db.DB.Where("id = ?", attachmentId).
		First(&attachment).Error
	if err != nil {
		return errors.New("attachment_not_found")
	}
	// 判断是否已经购买
	var exists bool
	err = db.DB.Model(&models.IntegralLog{}).
		Select("1").
		Where("user_id = ?", userId).
		Where("type = ?", models.IntegralTypeReduce).
		Where("source = ?", models.SourceTypePurchaseAttachment).
		Where("source_id = ?", attachmentId).
		Limit(1).
		Find(&exists).Error
	if err != nil {
		klog.Error("AttachmentBuy", err)
		return errors.New("unknown")
	}
	if exists {
		return errors.New("attachment_buy")
	}
	if attachment.Coins == 0 {
		return errors.New("attachment_free")
	}
	err = services.BuyAttachment(userId, uint(attachmentId), attachment.Coins)
	if err != nil {
		return err
	}
	return nil
}
