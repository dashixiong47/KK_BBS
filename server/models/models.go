package models

import (
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/utils/print"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

// gorm.Model 的定义
type Model struct {
	CreatedAt time.Time      `json:"created_at" gorm:"index:index_created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
type User struct {
	ID           uint          `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Username     string        `json:"username" gorm:"size:10;not null;unique;index:index_username"` // 用户名
	Password     string        `json:"password" gorm:"size:255;not null"`                            // 用户密码
	Email        string        `json:"email" gorm:"size:255;index:index_email"`                      // 用户邮箱
	Phone        string        `json:"phone" gorm:"size:11;"`                                        // 用户手机号
	Avatar       string        `json:"avatar" gorm:"size:255;"`                                      // 用户头像
	Nickname     string        `json:"nickname" gorm:"size:20;"`                                     // 用户昵称
	Introduction string        `json:"introduction" gorm:"size:255;"`                                // 用户简介
	Role         pq.Int64Array `json:"role" gorm:"type:integer[];"`                                  // 用户角色
	Status       int           `json:"status" gorm:"size:1;"`                                        // 用户状态
	Model
}

// Post 帖子
type Post struct {
	ID          uint          `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	UserID      uint          `json:"user_id" gorm:"not null;index:index_user_id"`      // 用户ID
	Title       string        `json:"title" gorm:"size:255;not null;index:index_title"` // 标题
	Tags        pq.Int64Array `json:"tags" gorm:"type:jsonb"`                           // 标签
	Covers      pq.Int64Array `json:"covers" gorm:"type:jsonb"`                         // 封面
	DefaultPost PostDefault   `gorm:"foreignKey:PostID"`
	VideoPost   PostVideo     `gorm:"foreignKey:PostID"`
	ImagePost   PostImage     `gorm:"foreignKey:PostID"`
	TextPost    PostText      `gorm:"foreignKey:PostID"`
	Model
}

// PostDefault 默认类型
type PostDefault struct {
	ID            uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	PostID        uint   `gorm:"unique;index"`                                               // 外键，并且是唯一的
	Content       string `json:"content" gorm:"type:text"`                                   // 内容
	Hidden        string `json:"hidden" gorm:"size:1;not null;default:0;index:index_hidden"` // 是否隐藏
	HiddenContent string `json:"hidden_content" gorm:"type:text"`                            // 隐藏内容

}

// PostVideo 视频
type PostVideo struct {
	ID           uint          `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	PostID       uint          `gorm:"unique;index"`                 // 外键，并且是唯一的
	VideoID      pq.Int64Array `json:"video_id" gorm:"type:jsonb"`   // 视频
	Introduction string        `json:"introduction" gorm:"size:255"` // 简介
}

// PostImage 图片
type PostImage struct {
	ID           uint          `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	PostID       uint          `gorm:"unique;index"`                 // 外键，并且是唯一的
	ImageID      pq.Int64Array `json:"image_id" gorm:"type:jsonb"`   // 图片
	Introduction string        `json:"introduction" gorm:"size:255"` // 简介
}

// PostText 文本
type PostText struct {
	ID           uint          `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	PostID       uint          `gorm:"unique;index"`                 // 外键，并且是唯一的
	TextID       pq.Int64Array `json:"text" gorm:"type:jsonb"`       // 文本
	Introduction string        `json:"introduction" gorm:"size:255"` // 简介
}

// Tag 标签
type Tag struct {
	ID   uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Name string `json:"name" gorm:"size:255;not null;unique;index:index_name"` // 标签名
	Model
}

type File struct {
	ID       uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Md5      string `json:"md5" gorm:"size:255;not null"`       // 文件md5
	FileName string `json:"file_name" gorm:"size:255;not null"` // 文件名
	Type     string `json:"type" gorm:"size:255;not null"`      // 文件类型
	Url      string `json:"url" gorm:"size:255;not null"`       // 文件地址
	Model
}

func init() {
	initData()
	autoMigrate()
	print.Info("Successfully connected to database!")
}
func initData() {
	if !db.DB.Migrator().HasTable(&User{}) {
		err := db.DB.Migrator().CreateTable(&User{})
		if err != nil {
			print.Error("Failed to connect to database: %v", err)
		}
		db.DB.Create(&User{
			Username: "admin",
			Password: "123456",
			Role:     pq.Int64Array{1},
		})
	}
}
func autoMigrate() {
	// 自动迁移
	err := db.DB.AutoMigrate(&Post{}, &PostDefault{}, &PostVideo{}, &PostImage{}, &PostText{}, &Tag{}, &File{})
	if err != nil {
		print.Info("Failed to connect to database: %v", err)
	}
}
