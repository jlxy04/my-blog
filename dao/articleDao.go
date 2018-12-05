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

func (dao ArticleDao) ListMain(topNum int) []models.BlogArticle {
	list := make([]models.BlogArticle, 0)
	err := dao.engine.Desc("create_time").Limit(topNum).Find(&list)
	if err != nil {
		log.Println(err)
	}
	return list
}

func (dao ArticleDao) ListTop(topNum int) []models.BlogArticle {
	list := make([]models.BlogArticle, 0)
	err := dao.engine.Desc("read").Limit(topNum).Find(&list)
	if err != nil {
		log.Println(err)
	}
	return list
}

func (dao ArticleDao) GetByIds(ids []string) []models.BlogArticle {
	list := make([]models.BlogArticle, 0)
	err := dao.engine.In("id", ids).Find(&list)
	if err != nil {
		log.Println(err)
	}
	return list
}

func (dao ArticleDao) GetById(id string) models.BlogArticle {
	article := models.BlogArticle{}
	bool, err := dao.engine.ID(id).Get(&article)
	if err != nil || !bool {
		log.Println(err, bool)
	}
	return article
}

func (dao ArticleDao) ListByCategoryId(categoryId string) []models.BlogArticle {
	list := make([]models.BlogArticle, 0)
	dao.engine.Where("category_id = ?", categoryId).Desc("create_time").Limit(10).Find(&list)
	return list
}
