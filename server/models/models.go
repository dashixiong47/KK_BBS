package models

import (
	"github.com/dashixiong47/KK_BBS/config"
	"time"

	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/server/data/group"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"gorm.io/gorm"
)

// Model gorm.Model 的定义
type Model struct {
	CreatedAt time.Time      `json:"created_at" gorm:"index:index_created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
type User struct {
	ID           uint         `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Username     string       `json:"username" gorm:"size:10;not null;unique;index:index_username"` // 用户名
	Password     string       `json:"password" gorm:"size:255;not null"`                            // 用户密码
	Email        string       `json:"email" gorm:"size:255;index:index_email"`                      // 用户邮箱
	Phone        string       `json:"phone" gorm:"size:11;"`                                        // 用户手机号
	Avatar       string       `json:"avatar" gorm:"size:255;"`                                      // 用户头像
	Nickname     string       `json:"nickname" gorm:"size:20;"`                                     // 用户昵称
	Background   string       `json:"background" gorm:"size:255;"`                                  // 用户背景
	Introduction string       `json:"introduction" gorm:"size:255;"`                                // 用户简介
	Role         *db.IntArray `json:"role" gorm:"type:integer[];"`                                  // 用户角色
	Status       int          `json:"status" gorm:"size:1;"`                                        // 用户状态
	Coins        int          `json:"coins" gorm:"-"`                                               // 金币
	Exp          int          `json:"exp" gorm:"-"`                                                 // 经验
	Level        int          `json:"level" gorm:"-"`                                               // 等级
	Follow       int          `json:"follow" gorm:"-"`                                              // 关注
	Fans         int          `json:"fans" gorm:"-"`                                                // 粉丝
	Model
}

// Topic 帖子
type Topic struct {
	ID         uint          `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	UserID     uint          `json:"-" gorm:"not null;index:index_user_id"`                               // 用户ID
	GroupID    uint          `json:"groupId" gorm:"not null;index:index_group_id" binding:"required"`     // 分组ID
	Title      string        `json:"title" gorm:"size:255;not null;index:index_title" binding:"required"` // 标题
	Tags       *db.IntArray  `json:"tags" gorm:"type:integer[];"`                                         // 标签
	Covers     *db.JSONSlice `json:"covers" gorm:"type:jsonb;"`                                           // 封面
	Type       int           `json:"type" gorm:"size:1;not null;index:index_type"`                        // 类型 1:默认 2:视频 3:图片 4:文本
	Summary    string        `json:"summary" gorm:"size:255;"`                                            // 摘要
	TopicBasic TopicBasic    `json:"topicBasic" gorm:"foreignKey:TopicID"`
	TopicVideo TopicVideo    `json:"topicVideo" gorm:"foreignKey:TopicID"`
	TopicImage TopicImage    `json:"topicImage" gorm:"foreignKey:TopicID"`
	TopicText  TopicText     `json:"topicText" gorm:"foreignKey:TopicID"`
	Model
}

// TopicBasic 默认类型
type TopicBasic struct {
	ID            uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	TopicID       uint   `gorm:"unique;index"`                                               // 外键，并且是唯一的
	Content       string `json:"content" gorm:"type:text"`                                   // 内容
	Hidden        int    `json:"hidden" gorm:"size:1;not null;default:0;index:index_hidden"` // 是否隐藏
	HiddenContent string `json:"hiddenContent" gorm:"type:text"`                             // 隐藏内容
}

// TopicVideo 视频
type TopicVideo struct {
	ID           uint         `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	TopicID      uint         `gorm:"unique;index"`                   // 外键，并且是唯一的
	VideoID      *db.IntArray `json:"videoId" gorm:"type:integer[];"` // 视频
	Introduction string       `json:"introduction" gorm:"size:255"`   // 简介
}

// TopicImage 图片
type TopicImage struct {
	ID           uint         `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	TopicID      uint         `gorm:"unique;index"`                   // 外键，并且是唯一的
	ImageID      *db.IntArray `json:"imageId" gorm:"type:integer[];"` // 图片
	Introduction string       `json:"introduction" gorm:"size:255"`   // 简介
}

// TopicText 文本
type TopicText struct {
	ID           uint         `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	TopicID      uint         `gorm:"unique;index"`                 // 外键，并且是唯一的
	TextID       *db.IntArray `json:"text" gorm:"type:integer[];"`  // 文本
	Introduction string       `json:"introduction" gorm:"size:255"` // 简介
}

// TopicView 浏览
type TopicView struct {
	ID      uint `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	UserID  uint `json:"userId" gorm:"not null;index:index_user_id"`   // 用户ID
	TopicID uint `json:"topicId" gorm:"not null;index:index_topic_id"` // 帖子ID
	Model
}

// Tag 标签
type Tag struct {
	ID   uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name string `json:"name" gorm:"size:255;not null;unique;index:index_name"` // 标签名
	Model
}

// File 文件
type File struct {
	ID       uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Md5      string `json:"md5" gorm:"size:255;not null;index:index_md5"` // 文件md5
	FileName string `json:"file_name" gorm:"size:255;not null"`           // 文件名
	Type     string `json:"type" gorm:"size:255;not null"`                // 文件类型
	Url      string `json:"url" gorm:"size:255;not null"`                 // 文件地址
	Model
}

// Group 分组
type Group struct {
	ID   uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name string `json:"name" gorm:"size:10;not null;unique;index:index_name"` // 分组名
	Icon string `json:"icon" gorm:"size:255;not null"`                        // 分组图标
	Model
}

// TopicLike 点赞
type TopicLike struct {
	ID      uint `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	UserID  uint `json:"userId" gorm:"not null;index:index_user_id"`   // 用户ID
	TopicID uint `json:"topicId" gorm:"not null;index:index_topic_id"` // 帖子ID
	Model
}

// Comment 评论
type Comment struct {
	ID            uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	ParentID      uint   `json:"parentId" gorm:"not null;index:index_parent_id"`    // 父级ID
	UserID        uint   `json:"userId" gorm:"not null;index:index_user_id"`        // 用户ID
	TopicID       uint   `json:"topicId" gorm:"not null;index:index_topic_id"`      // 帖子ID
	ReplyToUserID uint   `json:"replyToUserId" gorm:"index:index_reply_to_user_id"` // 被回复用户的ID
	Content       string `json:"content" gorm:"type:text;not null"`                 // 评论内容
	Model
}

// CommentLike 评论点赞
type CommentLike struct {
	ID           uint `gorm:"primaryKey;AUTO_INCREMENT"`
	UserID       uint `gorm:"primaryKey;index:index_user_id"`
	CommentID    uint `gorm:"primaryKey;index:index_comment_id"`
	SubCommentID uint `gorm:"primaryKey;index:index_sub_comment_id;default:0"`
	Model
}

// Attachment 附件
type Attachment struct {
	ID      uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	UserID  uint   `json:"userId" gorm:"not null;index:index_user_id"`               // 用户ID
	TopicID uint   `json:"topicId" gorm:"not null;index:index_topic_id"`             // 帖子ID
	Type    int    `json:"type" gorm:"size:10;not null"`                             // 文件类型 1文件 2网盘
	Coins   int    `json:"coins" gorm:"size:255;default:0"`                          // 金币
	Hidden  int    `json:"hidden" gorm:"size:1;not null;default:0"`                  // 是否隐藏
	Order   int    `json:"order" gorm:"size:1;not null;default:0;index:index_order"` // 排序
	Name    string `json:"name" gorm:"size:255;not null"`                            // 文件名

	FileType    string `json:"fileType" gorm:"size:255"`    // 文件类型
	FileSize    int    `json:"fileSize" gorm:"size:255"`    // 文件大小
	FileUrl     string `json:"url" gorm:"size:255"`         // 地址
	NetDiskType string `json:"netDisk" gorm:"size:255"`     // 网盘类型
	NetDiskUrl  string `json:"netDiskUrl" gorm:"size:255"`  // 网盘地址
	NetDiskCode string `json:"netDiskCode" gorm:"size:255"` // 网盘提取码
	Remake      string `json:"remake" gorm:"size:255"`      // 备注
	Model
}

// Collection 收藏
type Collection struct {
	ID      uint `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	UserID  uint `json:"userId" gorm:"not null;index:index_user_id"`   // 用户ID
	TopicID uint `json:"topicId" gorm:"not null;index:index_topic_id"` // 帖子ID
	Model
}

// Follow 关注
type Follow struct {
	ID       uint `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	UserID   uint `json:"userId" gorm:"not null;index:index_user_id"`     // 用户ID
	FollowID uint `json:"followId" gorm:"not null;index:index_follow_id"` // 关注ID
	Model
}

// Share 分享
type Share struct {
	ID      uint `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	UserID  uint `json:"userId" gorm:"not null;index:index_user_id"`   // 用户ID
	TopicID uint `json:"topicId" gorm:"not null;index:index_topic_id"` // 帖子ID
	Model
}

// IntegralLog Source 1:发帖 2:评论 3:被评论 4:签到 5:充值 6:购买附件 7:打赏 8:悬赏
// IntegralLog 积分系统 积分=金币
type IntegralLog struct {
	ID       uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	UserID   uint   `json:"userId" gorm:"not null;index:index_user_id"`   // 用户ID
	Type     int    `json:"type" gorm:"size:1;not null;index:index_type"` // 类型 1:增加 2:减少
	Source   int    `json:"source" gorm:"size:1;not null;index:index_source"`
	SourceID uint   `json:"sourceId" gorm:"not null;index:index_source_id"`
	Number   int    `json:"number" gorm:"size:255;default:0;"` // 数量
	Remake   string `json:"remake" gorm:"size:255;"`           // 备注
	Model
}

// RedeemCode 卡密密钥
type RedeemCode struct {
	ID     uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Code   string `json:"code" gorm:"size:255;not null;unique;index:index_code"` // 密钥
	Type   int    `json:"type" gorm:"size:1;not null;index:index_type"`          // 类型 1:金币 2:会员
	Number int    `json:"number" gorm:"size:255;default:0;"`                     // 数量 1:金币 2:会员 时长为天数
	Remake string `json:"remake" gorm:"size:255;"`                               // 备注
	Model
}

// Config admin 管理后台
// 配置
type Config struct {
	ID    uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Code  string `json:"code" gorm:"size:255;not null;unique;index:index_code"` // 配置代码
	Name  string `json:"name" gorm:"size:255;not null;unique;index:index_name"` // 配置名
	Value string `json:"value" gorm:"size:255;not null;"`                       // 配置值
	Model
}

func init() {
	initData()
	autoMigrate()
	LoadSystemConfig()
	klog.Info("Successfully connected to database!")
}
func initData() {
	if !db.DB.Migrator().HasTable(&User{}) {
		err := db.DB.Migrator().CreateTable(&User{})
		if err != nil {
			klog.Error("Failed to connect to database: %v", err)
		}
		db.DB.Create(&User{
			Username: "admin",
			Password: "e10adc3949ba59abbe56e057f20f883e", //123456
			Nickname: "admin",
			Role:     &db.IntArray{1, 2, 3},
		})
	}
	if !db.DB.Migrator().HasTable(&Group{}) {
		err := db.DB.Migrator().CreateTable(&Group{})
		if err != nil {
			klog.Error("Failed to connect to database: %v", err)
		}
		var groups = Group{
			Name: "默认分组",
		}
		db.DB.Create(&groups)
		err = group.CreateGroup(groups.ID, groups.Name)
		if err != nil {
			klog.Error("Failed to connect to database: %v", err)
		}
	}
}
func autoMigrate() {
	// 自动迁移
	err := db.DB.AutoMigrate(&TopicLike{}, &Topic{}, &TopicBasic{}, &TopicImage{}, &TopicVideo{}, &TopicText{}, &TopicView{}, &Tag{}, &File{}, &Comment{}, &CommentLike{},
		&Attachment{}, &Collection{}, &Follow{}, &Share{}, &IntegralLog{}, &RedeemCode{}, &Config{})
	if err != nil {
		klog.Info("Failed to connect to database: %v", err)
	}
}

// LoadSystemConfig 加载systemConfig
func LoadSystemConfig() {
	//	读取数据库里的配置
	var sysConfig []Config
	tx := db.DB.Find(&sysConfig)
	if tx.Error != nil {
		klog.Error("Error reading config: %s\n", tx.Error)
	}
	//	写到SystemConfig
	for _, v := range sysConfig {
		config.SystemConfig[v.Code] = v.Value
	}
	//	打印加载的配置
	klog.Info("successfullyLoaded:%v", config.SystemConfig)
}
