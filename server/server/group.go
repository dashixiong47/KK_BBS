package server

import (
	"github.com/dashixiong47/KK_BBS/db"
	"github.com/dashixiong47/KK_BBS/models"
	"github.com/dashixiong47/KK_BBS/utils/klog"
)

type GroupServer struct{}

func (g *GroupServer) GroupInfo() (any, error) {
	var groupInfo []models.Group
	err := db.DB.Find(&groupInfo).Error
	if err != nil {
		klog.Error(err.Error())
		return nil, err
	}
	return groupInfo, nil
}
