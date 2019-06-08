package controllers

import (
	"github.com/my-blog/models"
	"github.com/my-blog/service"
)

type MyDescController struct {
	Server service.MyDescService
}

func (c MyDescController) GetAll() []models.BlogMyDesc {
	return c.Server.GetAll()
}
