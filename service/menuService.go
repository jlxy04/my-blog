package service

import (
	"my-blog/dao"
	"my-blog/datasource"
	"my-blog/models"
)

type MenuService interface {
	GetAll() []models.BlogMenu

	CreateMenu(menu *models.BlogMenu) error

	GetById(id string) *models.BlogMenu
}

func NewMenuService() MenuService {
	return &menuService{
		dao: dao.NewMenuDao(datasource.InstanceMaster()),
	}
}

type menuService struct {
	dao *dao.MenuDao
}

func (s menuService) GetAll() []models.BlogMenu {
	return s.dao.GetAll()
}

func (s menuService) CreateMenu(menu *models.BlogMenu) error {
	return s.dao.Create(menu)
}

func (s menuService) GetById(id string) *models.BlogMenu {
	return s.dao.Get(id)
}
