package controllers

import (
	ArticleRepository "blog/internal/modules/article/repositories"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	//repositories.ArticleRepositoryInterface `inject:"articleRepositoryInterface"`
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *Controller {
	return &Controller{
		articleRepository: ArticleRepository.New(),
	}
}

func (ctrl *Controller) Index(c *gin.Context) {
	/*html.Render(c, http.StatusOK, "modules/home/view/home", gin.H{
		"title": "Home Page",
	})*/

	c.JSON(200, gin.H{
		"articles": ctrl.articleRepository.List(10),
	})
}
