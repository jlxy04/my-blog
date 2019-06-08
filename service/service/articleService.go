package service

import (
	"github.com/my-blog/dao"
	"github.com/my-blog/datasource"
	"github.com/my-blog/models"
)

type ArticleService interface {
	ListMain(topNum int) []models.BlogArticle

	ListTop(topNum int) []models.BlogArticle

	GetById(id string) models.BlogArticle

	ListByCategory(categoryId string) []models.BlogArticle
}

func NewArticleService() ArticleService {
	return &articleService{
		dao:           dao.NewArticleDao(datasource.InstanceMaster()),
		extendService: NewExtendService(),
	}
}

type articleService struct {
	dao           *dao.ArticleDao
	extendService ExtendService
}

func (s articleService) ListMain(topNum int) []models.BlogArticle {
	return s.dao.ListMain(topNum)
}

func (s articleService) ListTop(topNum int) []models.BlogArticle {
	ids := s.extendService.GetReadTopIds(topNum)
	return s.dao.GetByIds(ids)
}

//通过ID查询文章内容、阅读次数加1
func (s articleService) GetById(id string) models.BlogArticle {
	//阅读次数加1
	s.extendService.IncreaseReadCount(id, 1)
	return s.dao.GetById(id)
}

//通过ID查询文章内容、阅读次数加1
func (s articleService) ListByCategory(categoryId string) []models.BlogArticle {
	return s.dao.ListByCategoryId(categoryId)
}
