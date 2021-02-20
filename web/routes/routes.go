package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"superstar/bootstrap"
	"superstar/services"
	"superstar/web/controllers"
	"superstar/web/middleware"
)

func Configure(b *bootstrap.Bootstrapper) {
	superstarService := services.NewSuperstarService()

	index := mvc.New(b.Party("/"))
	index.Register(superstarService)
	index.Handle(new(controllers.AdminController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(superstarService)
	admin.Handle(new(controllers.AdminController))
}
