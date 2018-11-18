package service

import (
	"my-blog/dao"
	"my-blog/datasource"
	"my-blog/models"
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
