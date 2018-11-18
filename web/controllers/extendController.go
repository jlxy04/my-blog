package controllers

import (
	"my-blog/models"
	"my-blog/service"
)

type ExtendController struct {
	Service service.ExtendService
}

func (c ExtendController) GetReadtop() []models.BlogExtend {
	return nil
}
