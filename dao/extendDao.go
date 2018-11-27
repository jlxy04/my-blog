package dao

import (
	"github.com/go-xorm/xorm"
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
