package db

import (
	"fmt"
	"github.com/dashixiong47/KK_BBS/config"
	"github.com/dashixiong47/KK_BBS/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s", config.SettingsConfig.Postgresql.User, config.SettingsConfig.Postgresql.Password, config.SettingsConfig.Postgresql.Database, config.SettingsConfig.Postgresql.Host, config.SettingsConfig.Postgresql.Port, config.SettingsConfig.Postgresql.Sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.Error("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		utils.Error("Failed to connect to database: %v", err)
	}

	// 尝试与数据库建立连接
	if err := sqlDB.Ping(); err != nil {
		utils.Error("Failed to connect to database: %v", err)
	}
	utils.Info("Successfully connected to database!")
	DB = db
}
