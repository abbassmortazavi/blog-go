package services

import (
	"blog/internal/modules/article/models"
	"blog/internal/modules/article/repositories"
	"blog/internal/modules/article/responses"
)

type ArticleService struct {
	articleRepository repositories.ArticleRepositoryInterface
}

func New() *ArticleService {
	return &ArticleService{
		articleRepository: repositories.New(),
	}
}

func (s *ArticleService) GetFeaturedArticles() responses.Articles {
	articles := s.articleRepository.List(4)
	return responses.ToArticles(articles)
}
func (s *ArticleService) GetStoriesArticles() []models.Article {

	return s.articleRepository.List(6)
}
