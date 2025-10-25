package controllers

import "github.com/gin-gonic/gin"

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (ctrl *Controller) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World!",
	})
	//9:03
}
