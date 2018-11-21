package dao

import (
	"github.com/go-xorm/xorm"
	"my-blog/models"
)

type LeaveMsgDao struct {
	engine *xorm.Engine
}

func NewLeaveMsgDao(engine *xorm.Engine) *LeaveMsgDao {
	return &LeaveMsgDao{
		engine: engine,
	}
}

func (dao LeaveMsgDao) GetTopLeaveMsg(topNum int) []models.BlogLeaveMsg {
	list := make([]models.BlogLeaveMsg, 0)
	dao.engine.Desc("create_time").Limit(topNum).Find(&list)
	return list
}
