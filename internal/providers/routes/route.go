package routes

import (
	article "blog/internal/modules/article/routes"
	home "blog/internal/modules/home/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	home.Routes(router)
	article.Routes(router)
}
