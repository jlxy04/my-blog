package service

import (
	"my-blog/dao"
	"my-blog/datasource"
	"my-blog/models"
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
