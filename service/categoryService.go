package service

import (
	"github.com/my-blog/dao"
	"github.com/my-blog/datasource"
	"github.com/my-blog/models"
)

type CategoryService interface {
	GetMain() []models.BlogCategory
}

func NewCategoryService() CategoryService {
	return &categoryService{
		dao: dao.NewCategoryDao(datasource.InstanceMaster()),
	}
}

type categoryService struct {
	dao *dao.CategoryDao
}

func (s categoryService) GetMain() []models.BlogCategory {
	return s.dao.GetMain()
}
