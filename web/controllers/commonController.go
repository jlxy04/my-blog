package controllers

import (
	"my-blog/models"
	"my-blog/service"
)

type CommonController struct {
	Server service.CommonService
}

//根据文章ID查询评论
func (c CommonController) PostListArticleBy(articleId string) []models.BlogComment {
	return c.Server.ListCommonByArticleId(articleId)
}

func (c CommonController) PostAddCommon(name string, captcha string, content string) {

}
