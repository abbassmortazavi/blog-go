package bootstrap

import (
	"blog/pkg/config"
	"blog/pkg/database"
	"blog/pkg/html"
	"blog/pkg/routing"
	"blog/pkg/session"
	"blog/pkg/static"
)

func Serve() {
	config.Set()
	database.Connect()
	routing.Init()
	session.Start(routing.GetRouter())
	router := routing.GetRouter()
	static.LoadAsset(router)
	html.LoadHtml(router)
	routing.RegisterRoutes()
	routing.Run()
}
