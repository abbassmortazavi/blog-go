package routes

import (
	"blog/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/view/home", gin.H{
			"title": "Main website",
		})

	})

	router.GET("/about", func(c *gin.Context) {
		html.Render(c, http.StatusOK, "modules/home/view/about", gin.H{
			"title": "About  Page",
		})

	})
}
