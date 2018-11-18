package dao

import (
	"github.com/go-xorm/xorm"
	"log"
	"my-blog/models"
)

type CategoryDao struct {
	engine *xorm.Engine
}

func NewCategoryDao(engine *xorm.Engine) *CategoryDao {
	return &CategoryDao{
		engine: engine,
	}
}

func (d CategoryDao) GetMain() []models.BlogCategory {
	list := make([]models.BlogCategory, 0)
	err := d.engine.Where("show_main=?", "Y").Asc("sort").Find(&list)
	if err != nil {
		log.Println(err)
	}
	return list
}
