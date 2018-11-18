package controllers

import (
	"my-blog/models"
	"my-blog/service"
)

type ArticleController struct {
	Service service.ArticleService
}

func (c ArticleController) GetMain() []models.BlogArticle {
	return c.Service.ListMain(10)
}

func (c ArticleController) GetReadtop() []models.BlogArticle {
	return c.Service.ListTop(8)
}
