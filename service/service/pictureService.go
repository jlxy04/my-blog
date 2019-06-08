package service

import (
	"github.com/my-blog/dao"
	"github.com/my-blog/datasource"
	"github.com/my-blog/models"
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
