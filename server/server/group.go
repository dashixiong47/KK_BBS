package server

import (
	"github.com/dashixiong47/KK_BBS/server/data/group"
)

type GroupServer struct{}

func (g *GroupServer) GroupInfo() (any, error) {

	return group.GetAllGroup(), nil
}
