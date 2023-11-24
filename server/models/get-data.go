package models

import (
	"github.com/dashixiong47/KK_BBS/db"
)

// GetTopicType 获取Topic类型
func GetTopicType(_type int) string {
	var tableName string
	switch _type {
	case 1:
		tableName = TopicBasic{}.TableName()
	case 2:
		tableName = TopicImage{}.TableName()
	case 3:
		tableName = TopicVideo{}.TableName()
	case 4:
		tableName = TopicText{}.TableName()
	}
	return tableName
}

// GetUser 用户
func GetUser(id uint) (user User, err error) {
	err = db.DB.Model(User{}).Where("id = ?", id).First(&user).Error
	return
}
