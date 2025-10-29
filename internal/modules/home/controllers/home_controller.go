package controllers

import (
	ArticleRepository "blog/internal/modules/article/repositories"
	"blog/internal/modules/article/services"
	"blog/pkg/html"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	articleRepository ArticleRepository.ArticleRepositoryInterface
	articleService    services.ArticleServiceInterface
}

func New() *Controller {
	return &Controller{
		articleRepository: ArticleRepository.New(),
		articleService:    services.New(),
	}
}

func (ctrl *Controller) Index(c *gin.Context) {
	html.Render(c, http.StatusOK, "modules/home/view/home", gin.H{
		"title":    "Home Page",
		"featured": ctrl.articleService.GetFeaturedArticles(),
		"stories":  ctrl.articleService.GetFeaturedArticles(),
	})

	/*c.JSON(200, gin.H{
		"featured": ctrl.articleService.GetFeaturedArticles(),
		"stories":  ctrl.articleService.GetFeaturedArticles(),
	})*/
}
