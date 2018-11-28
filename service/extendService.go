package service

import (
	"my-blog/dao"
	"my-blog/datasource"
)

type ExtendService interface {
	GetReadTopIds(topNum int) []string

	IncreaseReadCount(articleId string, readCount int) (bool, error)

	IncreaseLikeCount(articleId string, readCount int) (bool, error)
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

//增加阅读次数
func (s extendService) IncreaseReadCount(articleId string, readCount int) (bool, error) {
	return s.dao.IncreaseReadCount(articleId, readCount)
}

// 增加喜欢次数
func (s extendService) IncreaseLikeCount(articleId string, readCount int) (bool, error) {
	return s.dao.IncreaseLikeCount(articleId, readCount)
}
