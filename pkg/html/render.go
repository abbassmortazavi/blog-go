package html

import (
	"blog/pkg/convertors"
	"blog/pkg/session"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Render(c *gin.Context, code int, name string, data gin.H) {
	data["app_name"] = viper.GetString("APPNAME")
	data["ERRORS"] = convertors.StringToMap(session.Flash(c, "errors"))
	c.HTML(code, name, data)
}
