package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/my-blog/service"
	"github.com/my-blog/web/controllers"
	"net/http"
	"strings"
	"time"
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

	// 全局session，默认30分钟
	session := sessions.New(sessions.Config{
		Expires: 30 * time.Minute,
	})

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

	mvc.Configure(app.Party("/meun"), func(app *mvc.Application) {
		menuService := service.NewMenuService()
		app.Register(menuService)
		app.Handle(new(controllers.MenuController))
	})

	mvc.Configure(app.Party("/my"), func(app *mvc.Application) {
		myDescService := service.NewMyDescService()
		app.Register(myDescService)
		app.Handle(new(controllers.MyDescController))
	})

	mvc.Configure(app.Party("/p"), func(app *mvc.Application) {
		pictureService := service.NewPictureService()
		app.Register(pictureService)
		app.Handle(new(controllers.PictureController))
	})

	mvc.Configure(app.Party("/c"), func(app *mvc.Application) {
		categoryService := service.NewCategoryService()
		app.Register(categoryService)
		app.Handle(new(controllers.CategoryController))
	})

	mvc.Configure(app.Party("/a"), func(app *mvc.Application) {
		articleService := service.NewArticleService()
		app.Register(articleService)
		app.Handle(new(controllers.ArticleController))
	})

	mvc.Configure(app.Party("/l"), func(app *mvc.Application) {
		//app.Router.Use(middleware.BasicAuth)
		leaveMsgService := service.NewLeaveMsgService()
		app.Register(leaveMsgService)
		app.Handle(new(controllers.LeaveMsgController))
	})

	mvc.Configure(app.Party("/e"), func(app *mvc.Application) {
		extendService := service.NewExtendService()
		app.Register(extendService)
		app.Handle(new(controllers.ExtendController))
	})

	mvc.Configure(app.Party("/co"), func(app *mvc.Application) {
		commonService := service.NewCommonService()
		app.Register(commonService)
		app.Register(session.Start, time.Now())
		app.Handle(new(controllers.CommonController))
	})

	mvc.Configure(app.Party("/ca"), func(app *mvc.Application) {
		//app.Router.Use(middleware.BasicAuth)
		captchaService := service.NewCaptchaService()
		app.Register(captchaService)
		app.Register(session.Start, time.Now())
		app.Handle(new(controllers.CaptchaController))
	})

	service.StartSchedules()

	app.Run(
		//开启web服务
		iris.Addr(":8080"),
		// 按下CTRL / CMD + C时跳过错误的服务器：
		iris.WithoutServerError(iris.ErrServerClosed),
		//实现更快的json序列化和更多优化：
		iris.WithOptimizations,
	)
}

type JsonTime time.Time

func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
