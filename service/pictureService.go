package service

import (
	"my-blog/dao"
	"my-blog/datasource"
	"my-blog/models"
)

type PictureService interface {
	GetTop(topNum int) []models.BlogPicture
}

type pictureService struct {
	dao *dao.PictureDao
}

func NewPictureService() PictureService {
	return &pictureService{
		dao: dao.NewPictureDao(datasource.InstanceMaster()),
	}
}

func (p pictureService) GetTop(topNum int) []models.BlogPicture {
	return p.dao.GetTop(topNum)
}
