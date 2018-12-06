package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/my-blog/extend"
	"github.com/my-blog/models"
	"github.com/my-blog/service"
)

type CommonController struct {
	Session *sessions.Session

	Server service.CommonService
}

//根据文章ID查询评论
func (c CommonController) PostListArticleBy(articleId string) []models.BlogComment {
	return c.Server.ListCommonByArticleId(articleId)
}

//增加评价
func (c CommonController) PostAddCommon(ctx iris.Context, session *sessions.Session) models.BlogComment {
	// 先校验验证码是否正确
	idKey := c.Session.GetString(ctx.FormValue("articleId"))
	rs := extend.VerfiyCaptcha(idKey, ctx.FormValue("captcha"))
	if rs {
		//如果验证码正确，则保存信息
		common := models.BlogComment{
			Id:          extend.GenerateId(),
			ArticleId:   ctx.FormValue("articleId"),
			Commentator: ctx.FormValue("name"),
			Content:     ctx.FormValue("content"),
		}
		c.Server.Create(common)
		return common
	}
	// 如果验证码不正确，则提示
	return models.BlogComment{}
}
