package db

import (
	"context"
	"fmt"
	"github.com/dashixiong47/KK_BBS/config"
	"github.com/dashixiong47/KK_BBS/utils"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

var Rdb *redis.Client

func init() {
	// 初始化 Redis 客户端，并设置连接池参数
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.SettingsConfig.Redis.Host, config.SettingsConfig.Redis.Port), // Redis 服务器地址
		Password: config.SettingsConfig.Redis.Password,                                                     // 密码（默认为空）
		DB:       config.SettingsConfig.Redis.Database,                                                     // 使用的数据库（默认为 0）

		// 连接池相关配置
		PoolSize:           config.SettingsConfig.Redis.PooSize,                                         // 连接池大小
		MinIdleConns:       config.SettingsConfig.Redis.MinIdleConns,                                    // 最小空闲连接数
		MaxConnAge:         time.Duration(config.SettingsConfig.Redis.MaxConnAge) * time.Minute,         // 连接的最大存活时间
		PoolTimeout:        time.Duration(config.SettingsConfig.Redis.PoolTimeout) * time.Second,        // 等待获取连接的最大时间
		IdleTimeout:        time.Duration(config.SettingsConfig.Redis.IdleTimeout) * time.Minute,        // 空闲连接的超时时间
		IdleCheckFrequency: time.Duration(config.SettingsConfig.Redis.IdleCheckFrequency) * time.Second, // 空闲连接检查频率
	})
	ctx := context.Background()
	// 检查连接
	ping, err := rdb.Ping(ctx).Result()
	if err != nil {
		klog.Error("连接 Redis 失败: %v", err)
		return
	}
	Rdb = rdb
	klog.Info("连接 Redis 成功: %s", ping)
}

// GetID 获取 Redis 中的值
func GetID(key interface{}) string {
	strID := fmt.Sprintf("%v", key)
	id, err := Rdb.HGet(context.Background(), "confusionID", strID).Result()
	if err == nil {
		return id
	}
	var encryptID string
	switch key.(type) {
	case int:
		encryptID = utils.EncryptID(key.(int))
		// 存储 key 和 encryptID，以及它们的反向键值对 uint同理
		Rdb.HSet(context.Background(), "confusionID", encryptID, key.(int), key.(int), encryptID)
	case uint:
		encryptID = utils.EncryptID(int(key.(uint)))
		Rdb.HSet(context.Background(), "confusionID", encryptID, key.(uint), key.(uint), encryptID)
	case string:
		decryptID, _ := utils.DecryptID(key.(string))
		Rdb.HSet(context.Background(), "confusionID", key.(string), decryptID, decryptID, key.(string))

	}

	return encryptID
}

// GetIntID 获取 Redis 中的值
func GetIntID(strID string) int {
	id, err := Rdb.HGet(context.Background(), "confusionID", strID).Result()
	if err == nil {
		// string 转 int
		idInt, _ := strconv.Atoi(id)
		return idInt
	}
	decryptID, _ := utils.DecryptID(strID)
	Rdb.HSet(context.Background(), "confusionID", strID, decryptID, decryptID, strID)

	return decryptID
}
