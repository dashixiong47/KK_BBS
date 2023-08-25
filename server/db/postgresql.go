package db

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dashixiong47/KK_BBS/config"
	"github.com/dashixiong47/KK_BBS/utils/print"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type JSONSlice []interface{}
type JSONMap []interface{}

func init() {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s", config.SettingsConfig.Postgresql.User, config.SettingsConfig.Postgresql.Password, config.SettingsConfig.Postgresql.Database, config.SettingsConfig.Postgresql.Host, config.SettingsConfig.Postgresql.Port, config.SettingsConfig.Postgresql.Sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print.Error("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		print.Error("Failed to connect to database: %v", err)
	}

	// 尝试与数据库建立连接
	if err := sqlDB.Ping(); err != nil {
		print.Error("Failed to connect to database: %v", err)
	}
	print.Info("Successfully connected to database!")
	//db = db.Debug()
	DB = db
}

func (j JSONSlice) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}
func (j JSONMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSONMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, j)
}
