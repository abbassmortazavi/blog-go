package bootstrap

import (
	"blog/pkg/config"
	"blog/pkg/routing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Serve() {
	config.Set()
	routing.Init()
	router := routing.GetRouter()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
			"app":     viper.GetString("APPNAME"),
		})
	})
	routing.Run()
}
