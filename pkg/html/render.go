package html

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Render(c *gin.Context, code int, name string, data gin.H) {
	data["app_name"] = viper.GetString("APPNAME")
	c.HTML(code, name, data)
}
