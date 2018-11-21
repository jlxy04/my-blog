package controllers

import (
	"my-blog/models"
	"my-blog/service"
)

type LeaveMsgController struct {
	Server service.LeaveMsgService
}

func (c LeaveMsgController) GetTop() []models.BlogLeaveMsg {
	return c.Server.GetTopLeaveMsg(10)
}
