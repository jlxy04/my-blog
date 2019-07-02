package service

import (
	"github.com/my-blog/dao"
	"github.com/my-blog/datasource"
	"github.com/my-blog/models"
)

type ExtendService interface {
	GetReadTopIds(topNum int) []string

	IncreaseReadCount(articleId string, readCount int) (bool, error)

	IncreaseLikeCount(articleId string, readCount int) (bool, error)

	GetByArticleId(articleId string) models.BlogExtend
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

//根据文章ID查询扩展详情
func (s extendService) GetByArticleId(articleId string) models.BlogExtend {
	return s.dao.GetByArticleId(articleId)
}
