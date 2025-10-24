package static

import "github.com/gin-gonic/gin"

func LoadAsset(router *gin.Engine) {
	router.Static("/assets", "./assets")
}
