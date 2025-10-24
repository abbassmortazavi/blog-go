package routes

import (
	home "blog/internal/modules/home/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	home.Routes(router)
}
