package controllers

import (
	userRepository "blog/internal/modules/user/repositories"
	userService "blog/internal/modules/user/services"
	"blog/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService    userService.UserServiceInterface
	userRepository userRepository.UserRepositoryInterface
}

func New() *Controller {
	return &Controller{
		userService:    userService.New(),
		userRepository: userRepository.New(),
	}
}

func (controller *Controller) Register(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/view/register", gin.H{
		"title": "Register",
	})
}

func (controller *Controller) HandelRegister(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "register success",
	})
}
