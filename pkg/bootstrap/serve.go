package bootstrap

import (
	"blog/pkg/config"
	"blog/pkg/html"
	"blog/pkg/routing"
)

func Serve() {
	config.Set()

	routing.Init()

	html.LoanHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
