package controllers

import (
	"my-blog/models"
	"my-blog/service"
)

type MyDescController struct {
	Server service.MyDescService
}

func (c MyDescController) GetAll() []models.BlogMyDesc {
	return c.Server.GetAll()
}
