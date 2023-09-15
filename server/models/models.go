package models

import (
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"gorm.io/gorm"
	"time"
)

// Model gorm.Model 的定义
type Model struct {
	CreatedAt time.Time      `json:"created_at" gorm:"index:index_created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
type User struct {
	ID           uint         `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Username     string       `json:"username" gorm:"size:10;not null;unique;index:index_username"` // 用户名
	Password     string       `json:"password" gorm:"size:255;not null"`                            // 用户密码
	Email        string       `json:"email" gorm:"size:255;index:index_email"`                      // 用户邮箱
	Phone        string       `json:"phone" gorm:"size:11;"`                                        // 用户手机号
	Avatar       string       `json:"avatar" gorm:"size:255;"`                                      // 用户头像
	Nickname     string       `json:"nickname" gorm:"size:20;"`                                     // 用户昵称
	Introduction string       `json:"introduction" gorm:"size:255;"`                                // 用户简介
	Role         *db.IntArray `json:"role" gorm:"type:integer[];"`                                  // 用户角色
	Status       int          `json:"status" gorm:"size:1;"`                                        // 用户状态
	Model
}

//// AfterFind 钩子，用于在查找之后解密 Email
//func (u *User) AfterFind(tx *gorm.DB) (err error) {
//	id, _ := utils.DecryptID(strconv.Itoa(int(u.ID)))
//	u.ID = uint(id)
//	return
//}

// Topic 帖子
type Topic struct {
	ID         uint          `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	UserID     uint          `json:"userId" gorm:"not null;index:index_user_id" binding:"required"`       // 用户ID
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

// Like 点赞
type Like struct {
	ID uint `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`

	Model
}

func init() {
	initData()
	autoMigrate()
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
			Role:     &db.IntArray{1, 2, 3},
		})
	}
	if !db.DB.Migrator().HasTable(&Group{}) {
		err := db.DB.Migrator().CreateTable(&Group{})
		if err != nil {
			klog.Error("Failed to connect to database: %v", err)
		}
		db.DB.Create(&Group{
			Name: "默认分组",
		})
	}
}
func autoMigrate() {
	// 自动迁移
	err := db.DB.AutoMigrate(&Topic{}, &TopicBasic{}, &TopicImage{}, &TopicVideo{}, &TopicText{}, &Tag{}, &File{})
	if err != nil {
		klog.Info("Failed to connect to database: %v", err)
	}
}
