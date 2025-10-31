package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Set(c *gin.Context, key string, value string) {
	session := sessions.Default(c)
	session.Set(key, value)
	err := session.Save()
	if err != nil {
		return
	}
}

func Flash(c *gin.Context, key string) string {
	session := sessions.Default(c)
	result := session.Get(key)
	session.Save()

	session.Delete(key)
	session.Save()

	if result != nil {
		return result.(string)
	}
	return ""
}
