package controllers

import (
	"my-blog/models"
	"my-blog/service"
)

type CategoryController struct {
	Server service.CategoryService
}

func (c CategoryController) GetMain() []models.BlogCategory {
	return c.Server.GetMain()
}
