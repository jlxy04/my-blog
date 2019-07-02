package service

import (
	"github.com/my-blog/dao"
	"github.com/my-blog/datasource"
	"github.com/my-blog/models"
)

type MyDescService interface {
	GetAll() []models.BlogMyDesc
}

func NewMyDescService() MyDescService {
	return &myDescService{
		dao: dao.NewMyDescDao(datasource.InstanceMaster()),
	}
}

type myDescService struct {
	dao *dao.MyDescDao
}

func (m myDescService) GetAll() []models.BlogMyDesc {
	return m.dao.GetAll()
}
