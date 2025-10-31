package controllers

import (
	userRepository "blog/internal/modules/user/repositories"
	"blog/internal/modules/user/requests"
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
	var request requests.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	user, err := controller.userService.Register(request)
	if err != nil {
		c.Redirect(http.StatusFound, "/register")
		return
	}
	/*html.Render(c, http.StatusOK, "modules/user/view/register", gin.H{
		"title": "Register",
	})*/

	c.JSON(200, gin.H{
		"message": "register success",
		"user":    user,
	})
}

func (controller *Controller) Login(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/user/view/login", gin.H{
		"title": "Login",
	})
}

func (controller *Controller) HandelLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login success",
	})
}
