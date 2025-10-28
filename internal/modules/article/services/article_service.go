package services

import (
	"blog/internal/modules/article/models"
	"blog/internal/modules/article/repositories"
)

type ArticleService struct {
	articleRepository repositories.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: repositories.New(),
	}
}

func (s *ArticleService) GetFeaturedArticles() []models.Article {
	return s.articleRepository.List(4)
}
func (s *ArticleService) GetStoriesArticles() []models.Article {
	return s.articleRepository.List(6)
}
