package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils/klog"
	"github.com/go-redis/redis/v8"
	"math"
	"time"
)

var hName = "userIntegral"
var maxRetries = 3

// AddIntegral 增加积分
// AddIntegral Source 1:发帖 2:评论 3:被评论 4:签到 5:充值 6:购买附件 7:打赏 8:悬赏
func AddIntegral(userId uint, _type, num, source int, sourceID uint, remake string) error {
	userIdStr := fmt.Sprintf("%v", userId)
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var err error
	if _type == 1 {
		err = performTransaction(db.Rdb, hName, userIdStr, int64(num))
	}
	if _type == 2 {
		err = performTransaction(db.Rdb, hName, userIdStr, int64(-num))
	}

	if err != nil {
		tx.Rollback()
		return err
	}
	var doc models.IntegralLog
	doc.UserID = userId
	doc.Type = _type
	doc.Number = num
	doc.Source = source
	doc.SourceID = sourceID
	doc.Remake = remake
	err = tx.Create(&doc).Error
	if err != nil {
		tx.Rollback()
		klog.Error("Create IntegralLog", err)
		return errors.New("unknown")
	}
	err = tx.Commit().Error

	return nil
}

var ctx = context.Background()

func performTransaction(rdb *redis.Client, key string, field string, change int64) error {
	script := `
    local currentBalance = tonumber(redis.call('HGET', KEYS[1], KEYS[2])) or 0
    local newBalance = currentBalance + tonumber(ARGV[1])
    if newBalance < 0 then
        return redis.error_reply('insufficient_funds')
    end
    redis.call('HSET', KEYS[1], KEYS[2], tostring(newBalance))
    return newBalance
    `
	var err error
	for i := 0; i < maxRetries; i++ {
		_, err = rdb.Eval(ctx, script, []string{key, field}, change).Result()
		if err != nil {
			if err == redis.TxFailedErr {
				// 指数退避策略
				time.Sleep(time.Millisecond * time.Duration(math.Pow(2, float64(i))))
				continue
			}
			klog.Error("performTransaction", err)
			// 积分更新失败
			return errors.New("insufficient_funds")
		}
		// 事务成功
		return nil
	}
	// 超过最大重试次数
	klog.Error("performTransaction_max_retries_exceeded", err)
	return errors.New("max_retries_exceeded")
}
