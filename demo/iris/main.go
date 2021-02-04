package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()

	htmlEngine := iris.HTML("./demo/iris/", ".html")
	app.RegisterView(htmlEngine)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("hello world -- from iris")
	})

	app.Get("/hello", func (ctx iris.Context) {
		ctx.ViewData("Title", "测试页面")
		ctx.ViewData("Content", "hello world -- from template")
		ctx.View("hello.html")
	})

	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}
