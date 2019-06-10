package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/context"
)

func main() {
	app := newApp()

	app.Run(iris.Addr(":8000"))
}

func newApp() *iris.Application {

	app := iris.New()

	//设置日志级别
	app.Logger().SetLevel("debug")

	//设置静态资源路径
	//app.StaticWeb("/static", "./web/html")
	//url：http://loccalhost:8000/web/html/hello.html
	//app.StaticWeb("/web/html", "./web/html/")
	app.StaticWeb("/static", "./web/html/")

	app.Get("/", func(context context.Context) {
		context.HTML("<h1> Iris 框架学习 </h1>")
	})

	return app
}

type VisistController struct {
	Session *sessions.Session
}

func (c *VisistController) Get() {
	c.Session.Set("visits", 1)
}
