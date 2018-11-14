package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.JSON(Test{Id: "xxxx", Age: 20})
	})
	app.Run(iris.Addr(":8080"))
}

type Test struct {
	Id  string
	Age int
}
