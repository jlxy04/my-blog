package controllers

import (
	"github.com/my-blog/models"
	"github.com/my-blog/service"
)

type CategoryController struct {
	Server service.CategoryService
}

func (c CategoryController) GetMain() []models.BlogCategory {
	return c.Server.GetMain()
}
