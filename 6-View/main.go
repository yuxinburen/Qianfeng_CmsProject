package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {

	//新建app对象
	app := iris.New()

	//注册视图文件
	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", func(context context.Context) {
		//为视图模板绑定数据
		context.ViewData("message", "hello world")
		//设置视图模板文件
		context.View("hello.html")
	})

	app.Run(iris.Addr(":8000"), iris.WithoutServerError(iris.ErrServerClosed))
}

