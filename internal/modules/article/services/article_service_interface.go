package services

import (
	"blog/internal/modules/article/models"
	"blog/internal/modules/article/responses"
)

type ArticleServiceInterface interface {
	GetFeaturedArticles() responses.Articles
	GetStoriesArticles() []models.Article
	Find(id int) (responses.Article, error)
}
