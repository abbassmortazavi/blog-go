package services

import (
	"blog/internal/modules/article/models"
	"blog/internal/modules/article/repositories"
	"blog/internal/modules/article/responses"
	"errors"
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

func (s *ArticleService) Find(id int) (responses.Article, error) {
	var response responses.Article
	article := s.articleRepository.Find(id)
	if article.ID == 0 {
		return response, errors.New("article not found")
	}
	return responses.ToArticle(article), nil
}
