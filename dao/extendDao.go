package dao

import (
	"github.com/go-xorm/xorm"
	"github.com/my-blog/models"
	"log"
)

type ExtendDao struct {
	engine *xorm.Engine
}

func NewExtendDao(engine *xorm.Engine) *ExtendDao {
	return &ExtendDao{
		engine: engine,
	}
}

func (e ExtendDao) GetReadTopIds(topNum int) []string {
	list := make([]string, 0)
	err := e.engine.Table("blog_extend").Cols("article_id").Desc("read_count").Limit(topNum).Find(&list)
	if err != nil {
		log.Println(err)
	}
	return list
}

func (dao ExtendDao) IncreaseReadCount(articleId string, readCount int) (bool, error) {
	sql := "update blog_extend set read_count = read_count + ? where article_id = ?"
	_, err := dao.engine.Exec(sql, readCount, articleId)
	return true, err
}

func (dao ExtendDao) IncreaseLikeCount(articleId string, readCount int) (bool, error) {
	sql := "update blog_extend set like_count = like_count + ? where article_id = ?"
	_, err := dao.engine.Exec(sql, readCount, articleId)
	return true, err
}

func (dao ExtendDao) GetByArticleId(articleId string) models.BlogExtend {
	list := make([]models.BlogExtend, 0)
	//sql := "select a.*, b.title as pre_article_name, c.title as next_article_name from blog_extend a left join blog_article b on a.pre_article = b.id left join blog_article c on a.nex_article = c.id where a.article_id = ?"
	dao.engine.Alias("a").Join("left", "blog_article b", "a.pre_article = b.id").
		Join("left", "blog_article c", "a.nex_article = c.id").
		Where("a.article_id = ?", articleId).Select("a.*, b.title as pre_article_name, c.title as next_article_name").Find(&list)
	if len(list) > 0 {
		return list[0]
	}
	return models.BlogExtend{}
}

func (dao ExtendDao) CountAll() models.BlogExtend {
	extend := models.BlogExtend{}
	dao.engine.Sums(&extend, "read_count", "like_count")
	return extend
}
