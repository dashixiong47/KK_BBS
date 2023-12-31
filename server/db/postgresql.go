package db

import (
	"fmt"
	"github.com/dashixiong47/KK_BBS/config"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s", config.SettingsConfig.Postgresql.User, config.SettingsConfig.Postgresql.Password, config.SettingsConfig.Postgresql.Database, config.SettingsConfig.Postgresql.Host, config.SettingsConfig.Postgresql.Port, config.SettingsConfig.Postgresql.Sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		//	打印出sql语句
		//Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "kk_", // 表名前缀
			SingularTable: true,  // 使用单数表名，禁用默认表名复数化
		},
	})
	if err != nil {
		klog.Error("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		klog.Error("Failed to connect to database: %v", err)
	}

	// 尝试与数据库建立连接
	if err := sqlDB.Ping(); err != nil {
		klog.Error("Failed to connect to database: %v", err)
	}
	klog.Info("Successfully connected to database!")
	//db = db.Debug()
	DB = db
}
