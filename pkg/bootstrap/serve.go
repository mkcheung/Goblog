package bootstrap

import (
	"blog/pkg/config"
	"blog/pkg/html"
	"blog/pkg/routing"
	"blog/pkg/static"
)

func Serve() {
	config.Set()

	routing.Init()

	static.LoadStatic(routing.GetRouter())

	html.LoanHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
