package controllers

import (
	"github.com/kataras/iris"
	"github.com/my-blog/extend"
	"github.com/my-blog/models"
	"github.com/my-blog/service"
	"time"
)

type LeaveMsgController struct {
	Server service.LeaveMsgService
}

func (c LeaveMsgController) GetTop() []models.BlogLeaveMsg {
	return c.Server.GetTopLeaveMsg(10)
}

func (c LeaveMsgController) PostAdd(ctx iris.Context) {
	//lm := &models.BlogLeaveMsg{
	//	Name:name,
	//	Mail:email,
	//	Content:content,
	//	CreateTime:extend.Time(time.Now()),
	//}
	msg := models.BlogLeaveMsg{}
	ctx.ReadForm(&msg)
	msg.Id = extend.GenerateId()
	msg.CreateTime = extend.Time(time.Now())
	c.Server.Create(msg)
}
