package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/my-blog/models"
	"log"
)

type PictureDao struct {
	engine *xorm.Engine
}

func NewPictureDao(engine *xorm.Engine) *PictureDao {
	return &PictureDao{
		engine: engine,
	}
}

func (p PictureDao) GetTop(topNum int) []models.BlogPicture {
	list := []models.BlogPicture{}
	err := p.engine.Asc("hot").Limit(topNum).Find(&list)
	if err != nil {
		log.Println(err)
		return list
	}
	return list
}
