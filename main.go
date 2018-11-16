package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"my-blog/service"
	"my-blog/web/controllers"
	"net/http"
	"strings"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.HTML("<b>Resource Not found</b>")
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.ServeFile("./html/index.html", false)
	})

	fileServer := app.StaticHandler("./html", false, false)

	app.WrapRouter(func(w http.ResponseWriter, r *http.Request, router http.HandlerFunc) {
		path := r.URL.Path
		if !strings.Contains(path, ".") {
			router(w, r)
			return
		}
		ctx := app.ContextPool.Acquire(w, r)
		fileServer(ctx)
		app.ContextPool.Release(ctx)
	})

	// 跨域
	//app.WrapRouter(func(w http.ResponseWriter, r *http.Request, firstNextIsTheRouter http.HandlerFunc) {
	//	w.Header().Add("Access-Control-Allow-Origin", "*")
	//	w.Header().Add("Access-Control-Allow-Credentials", "true")
	//	w.Header().Add("Access-Control-Allow-Methods", "GET,POST")
	//	w.Header().Add("Access-Control-Allow-Headers", "Set-Cookie,X-Requested-With,Content-Type")
	//})
	//app.Get("/", func(ctx iris.Context) {
	//	ctx.WriteString("欢迎....")
	//})

	//app.StaticWeb("/js", "./html/js")
	//app.StaticWeb("/css", "./html/css")
	//app.StaticWeb("/html", "./html")
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
