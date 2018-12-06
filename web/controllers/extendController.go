package controllers

import (
	"github.com/my-blog/models"
	"github.com/my-blog/service"
)

type ExtendController struct {
	Service service.ExtendService
}

func (c ExtendController) GetReadtop() []models.BlogExtend {
	return nil
}

func (c ExtendController) GetArticleBy(articleId string) models.BlogExtend {
	return c.Service.GetByArticleId(articleId)
}

func (c ExtendController) GetAddlikeBy(articleId string) models.BlogExtend {
	c.Service.IncreaseLikeCount(articleId, 1)
	return c.GetArticleBy(articleId)
}
