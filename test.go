package main

import (
	"fmt"
	"my-blog/extend"
)

func main() {
	idKeyD, base64stringD := extend.GenerateCaptcha()
	fmt.Println(base64stringD)
	extend.VerfiyCaptcha(idKeyD, idKeyD)
}
