package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"my-blog/service"
	"my-blog/web/controllers"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.StaticWeb("/js", "./html/js")
	app.StaticWeb("/css", "./html/css")
	app.StaticWeb("/html", "./html")
	mvc.Configure(app.Party("/meun"), menu)
	app.Run(
		//开启web服务
		iris.Addr(":8080"),
		// 按下CTRL / CMD + C时跳过错误的服务器：
		iris.WithoutServerError(iris.ErrServerClosed),
		//实现更快的json序列化和更多优化：
		iris.WithOptimizations,
	)
}

func menu(app *mvc.Application) {
	//app.Router.Use(middleware.BasicAuth)
	menuService := service.NewMenuService()
	app.Register(menuService)
	app.Handle(new(controllers.MenuController))
}
