package bootstrap

import (
	"blog/pkg/config"
	"blog/pkg/database"
	"blog/pkg/html"
	"blog/pkg/routing"
	"blog/pkg/static"
)

func Serve() {
	config.Set()
	database.Connect()
	routing.Init()
	router := routing.GetRouter()
	static.LoadAsset(router)
	html.LoadHtml(router)
	routing.RegisterRoutes()
	routing.Run()
}
