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
	router := routing.GetRouter()
	static.LoadAsset(router)
	html.LoadHtml(router)
	routing.RegisterRoutes()
	routing.Run()
}
