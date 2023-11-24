package models

import (
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/utils/klog"
)

// SysConfig admin 管理后台
// 配置
type SysConfig struct {
	ID    uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Code  string `json:"code" gorm:"size:255;not null;unique;index:index_code"` // 配置代码
	Name  string `json:"name" gorm:"size:255;not null;unique;index:index_name"` // 配置名
	Value string `json:"value" gorm:"size:255;not null;"`                       // 配置值
	Model
}

// SysUser admin 管理后台
// 用户
type SysUser struct {
	ID       uint   `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Username string `json:"username" gorm:"size:255;not null;unique;index:index_username"` // 用户名
	Password string `json:"password" gorm:"size:255;not null;"`                            // 密码
	Nickname string `json:"nickname" gorm:"size:255;not null;"`                            // 昵称
	Avatar   string `json:"avatar" gorm:"size:255;not null;"`                              // 头像
	Email    string `json:"email" gorm:"size:255;not null;"`                               // 邮箱
	Phone    string `json:"phone" gorm:"size:255;not null;"`                               // 手机号
	Role     string `json:"role" gorm:"size:255;not null;"`                                // 角色
	Model
}

func init() {
	initAdminData()
}
func initAdminData() {

	// 初始化用户
	if !db.DB.Migrator().HasTable(&SysUser{}) {
		err := db.DB.Migrator().CreateTable(&SysUser{})
		if err != nil {
			klog.Error("Failed to connect to database: %v", err)
		}
		err = db.DB.Create(&SysUser{
			Username: "admin",
			Password: "e10adc3949ba59abbe56e057f20f883e", //123456
			Nickname: "admin",
		}).Error
	}
	if !db.DB.Migrator().HasTable(&SysConfig{}) {

		err := db.DB.Migrator().CreateTable(&SysConfig{})
		if err != nil {
			klog.Error("Failed to connect to database: %v", err)
		}
		var _config = []SysConfig{
			{
				Code:  "site_name",
				Name:  "站点名称",
				Value: "KK_BBS",
			},
			{
				Code:  "site_host",
				Name:  "站点地址",
				Value: "http://localhost:8080/",
			},
		}
		//	创建默认配置
		err = db.DB.Create(&_config).Error
		if err != nil {
			klog.Error("Failed to connect to database: %v", err)
		}
	}
}
