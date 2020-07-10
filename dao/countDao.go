package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/my-blog/models"
)

func NewCountDao(engine *xorm.Engine) *CountDao {
	return &CountDao{
		engine: engine,
	}
}

type CountDao struct {
	engine *xorm.Engine
}

func (dao CountDao) GetAll() (list []models.BlogCount) {
	list = make([]models.BlogCount, 5)
	dao.engine.AllCols().Get(&list)
	return
}

func (dao CountDao) UpdateData(list []models.BlogCount) {
	session := dao.engine.NewSession()
	defer session.Clone()

	session.Begin()

	var err error
	for _, m := range list {
		_, err = session.Update(&m)
		if err != nil {
			break
		}
	}

	if err != nil {
		session.Rollback()
	}

	session.Commit()
}
