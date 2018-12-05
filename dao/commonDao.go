package dao

import (
	"github.com/go-xorm/xorm"
	"my-blog/models"
)

type CommonDao struct {
	engine *xorm.Engine
}

func NewCommonDao(engine *xorm.Engine) *CommonDao {
	return &CommonDao{
		engine: engine,
	}
}

func (dao CommonDao) ListCommonByArticleId(articleId string, topNum int) []models.BlogComment {
	list := make([]models.BlogComment, 0)
	dao.engine.Where("article_id = ?", articleId).Desc("create_time").Limit(10).Find(&list)
	return list
}

func (dao CommonDao) Create(comment models.BlogComment) {
	dao.engine.Insert(comment)
}
