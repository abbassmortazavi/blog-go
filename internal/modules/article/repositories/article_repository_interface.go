package repositories

import (
	"blog/internal/modules/article/models"
)

type ArticleRepositoryInterface interface {
	List(limit int) []models.Article
	Find(id int) models.Article
}
