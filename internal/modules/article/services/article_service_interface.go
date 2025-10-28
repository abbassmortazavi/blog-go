package services

import "blog/internal/modules/article/models"

type ArticleServiceInterface interface {
	GetFeaturedArticles() []models.Article
	GetStoriesArticles() []models.Article
}
