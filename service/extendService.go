package service

import (
	"my-blog/dao"
	"my-blog/datasource"
)

type ExtendService interface {
	GetReadTopIds(topNum int) []string
}

func NewExtendService() ExtendService {
	return &extendService{
		dao: dao.NewExtendDao(datasource.InstanceMaster()),
	}
}

type extendService struct {
	dao *dao.ExtendDao
}

func (s extendService) GetReadTopIds(topNum int) []string {
	return s.dao.GetReadTopIds(topNum)
}
