package controllers

import (
	"github.com/kataras/iris/sessions"
	"github.com/my-blog/service"
)

type CaptchaController struct {
	Session *sessions.Session

	Service service.CaptchaService
}

//生成验证码
func (c CaptchaController) GetGenCaptchaBy(articleId string) string {
	idKey, base64Str := c.Service.GeneratorCaptchaCode()
	c.Session.Set(articleId, idKey)
	return base64Str
}
