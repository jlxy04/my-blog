package dao

import (
	"github.com/go-xorm/xorm"
	"log"
	"my-blog/models"
)

type MyDescDao struct {
	engine *xorm.Engine
}

func NewMyDescDao(engine *xorm.Engine) *MyDescDao {
	return &MyDescDao{
		engine: engine,
	}
}

func (m MyDescDao) GetAll() []models.BlogMyDesc {
	list := make([]models.BlogMyDesc, 0)
	err := m.engine.Find(&list)
	if err != nil {
		log.Println(err)
	}
	return list
}
