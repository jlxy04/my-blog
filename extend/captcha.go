package extend

import (
	"github.com/mojocn/base64Captcha"
	"log"
)

//生成验证码，返回验证码和base64字符串
//func GenerateCaptcha() (string, string) {
//	//数字验证码配置
//	var configC = base64Captcha.ConfigCharacter{
//		Height: 40,
//		Width:  100,
//		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
//		Mode:               base64Captcha.CaptchaModeNumber,
//		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
//		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
//		IsUseSimpleFont:    true,
//		IsShowHollowLine:   true,
//		IsShowNoiseDot:     false,
//		IsShowNoiseText:    false,
//		IsShowSlimeLine:    false,
//		IsShowSineLine:     false,
//		CaptchaLen:         4,
//	}
//
//	idKey, cap := base64Captcha.GenerateCaptcha("", configC)
//
//	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(cap)
//
//	return idKey, base64stringD
//}
var store = base64Captcha.DefaultMemStore

func GenerateCaptcha() (string, string) {
	math := base64Captcha.DriverMath{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 0,
		BgColor:         nil,
		Fonts:           nil,
	}
	driver := math.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	if err != nil {
		log.Fatal(err)
	}
	return id, b64s
}

func VerfiyCaptcha(idkey, verifyValue string) (result bool) {
	result = store.Verify(idkey, verifyValue, true)
	return
}
