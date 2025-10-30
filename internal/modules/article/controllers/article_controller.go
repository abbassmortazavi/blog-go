package controllers

import (
	ArticleRepository "blog/internal/modules/article/repositories"
	ArticleService "blog/internal/modules/article/services"
	"blog/pkg/html"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleService    ArticleService.ArticleServiceInterface
	articleRepository ArticleRepository.ArticleRepositoryInterface
}

func New() *Controller {
	return &Controller{
		articleRepository: ArticleRepository.New(),
		articleService:    ArticleService.New(),
	}
}

func (controller *Controller) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		html.Render(c, http.StatusInternalServerError, "templates/html/errors/500", gin.H{
			"title": "Error",
			"error": err.Error(),
		})
		return
	}

	article, err := controller.articleService.Find(id)
	if err != nil {
		html.Render(c, http.StatusNotFound, "templates/html/errors/404", gin.H{
			"title": "Error",
			"error": err.Error(),
		})
		return
	}
	html.Render(c, http.StatusOK, "modules/article/view/show", gin.H{
		"title":   "Error",
		"article": article,
	})

}
