package dao

import (
	"github.com/go-xorm/xorm"
	"log"
	"my-blog/models"
)

type ArticleDao struct {
	engine *xorm.Engine
}

func NewArticleDao(engine *xorm.Engine) *ArticleDao {
	return &ArticleDao{
		engine: engine,
	}
}

func (a ArticleDao) ListMain(topNum int) []models.BlogArticle {
	list := make([]models.BlogArticle, 0)
	err := a.engine.Desc("create_time").Limit(topNum).Find(&list)
	if err != nil {
		log.Println(err)
	}
	return list
}

func (a ArticleDao) ListTop(topNum int) []models.BlogArticle {
	list := make([]models.BlogArticle, 0)
	err := a.engine.Desc("read").Limit(topNum).Find(&list)
	if err != nil {
		log.Println(err)
	}
	return list
}

func (a ArticleDao) GetByIds(ids []string) []models.BlogArticle {
	list := make([]models.BlogArticle, 0)
	err := a.engine.In("id", ids).Find(&list)
	if err != nil {
		log.Println(err)
	}
	return list
}
