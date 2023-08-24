package models

import (
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/utils"
	"gorm.io/gorm"
	"log"
	"time"
)

// gorm.Model 的定义
type Model struct {
	ID        uint `gorm:"primaryKey;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type User struct {
	Model
	Username     string `json:"username" gorm:"size:10;"`      // 用户名
	Password     string `json:"password" gorm:"size:10;"`      // 用户密码
	Email        string `json:"email" gorm:"size:255;"`        // 用户邮箱
	Phone        string `json:"phone" gorm:"size:11;"`         // 用户手机号
	Avatar       string `json:"avatar" gorm:"size:255;"`       // 用户头像
	Nickname     string `json:"nickname" gorm:"size:20;"`      // 用户昵称
	Introduction string `json:"introduction" gorm:"size:255;"` // 用户简介
	Role         int    `json:"role" gorm:"size:1;"`           // 用户角色
	Status       int    `json:"status" gorm:"size:1;"`         // 用户状态
}

func init() {
	log.Println("--------------------1111")
	// 自动迁移
	err := db.DB.AutoMigrate(&User{})
	if err != nil {
		utils.Info("Failed to connect to database: %v", err)
	}
}
