package service

import "github.com/my-blog/extend"

type CaptchaService interface {
	GeneratorCaptchaCode() (string, string)
}

func NewCaptchaService() CaptchaService {
	return &captchaService{}
}

type captchaService struct{}

//生成验证码图片
func (s captchaService) GeneratorCaptchaCode() (string, string) {
	return extend.GenerateCaptcha()
}
