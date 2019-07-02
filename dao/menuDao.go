package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/my-blog/models"
	"log"
)

type MenuDao struct {
	engine *xorm.Engine
}

func NewMenuDao(engine *xorm.Engine) *MenuDao {
	return &MenuDao{
		engine: engine,
	}
}

func (dao MenuDao) Get(id string) *models.BlogMenu {
	data := &models.BlogMenu{Id: id}
	ok, err := dao.engine.Get(data)
	if ok && err != nil {
		return data
	} else {
		data.Id = ""
		return data
	}
}

func (dao MenuDao) GetAll() []models.BlogMenu {
	datalist := []models.BlogMenu{}
	err := dao.engine.Desc("sort").Where("status = ?", "Y").Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	} else {
		return datalist
	}
}

func (dao MenuDao) Update(menu models.BlogMenu, colums []string) error {
	_, err := dao.engine.Id(menu.Id).MustCols(colums...).Update(menu)
	return err
}

func (dao MenuDao) Delete(id string) error {
	data := &models.BlogMenu{Id: id, Status: "N"}
	_, err := dao.engine.Id(data.Id).Update(data)
	return err
}

func (dao MenuDao) Create(menu *models.BlogMenu) error {
	_, err := dao.engine.Insert(menu)
	return err
}
