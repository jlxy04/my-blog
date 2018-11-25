package service

import (
	"my-blog/dao"
	"my-blog/datasource"
	"my-blog/models"
)

type LeaveMsgService interface {
	GetTopLeaveMsg(topNum int) []models.BlogLeaveMsg
	Create(msg models.BlogLeaveMsg)
}

func NewLeaveMsgService() LeaveMsgService {
	return &leaveMsgService{
		dao: dao.NewLeaveMsgDao(datasource.InstanceMaster()),
	}
}

type leaveMsgService struct {
	dao *dao.LeaveMsgDao
}

func (s leaveMsgService) GetTopLeaveMsg(topNum int) []models.BlogLeaveMsg {
	return s.dao.GetTopLeaveMsg(topNum)
}

func (s leaveMsgService) Create(msg models.BlogLeaveMsg) {
	s.dao.Create(msg)
}
