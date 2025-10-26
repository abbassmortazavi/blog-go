package controllers

import (
	"blog/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (ctrl *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/view/home", gin.H{
		"title": "Home Page",
	})
}
