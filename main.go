package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	app.Use(recover2.New())
	app.Use(logger.New())

	htmlEngine := iris.HTML("./templates", ".html")
	app.RegisterView(htmlEngine)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("welcome")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.ViewData("Title", "这是个标题")
		ctx.ViewData("Content", "这是内容")
		ctx.View("hello.html")
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
