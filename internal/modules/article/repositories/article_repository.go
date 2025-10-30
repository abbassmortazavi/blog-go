package repositories

import (
	"blog/internal/modules/article/models"
	"blog/pkg/database"

	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func New() *ArticleRepository {
	return &ArticleRepository{
		DB: database.Connection(),
	}
}
func (r *ArticleRepository) List(limit int) []models.Article {
	var articles []models.Article
	r.DB.Limit(limit).Joins("User").Order("rand()").Find(&articles)
	return articles
}

func (r *ArticleRepository) Find(id int) models.Article {
	var article models.Article
	r.DB.Joins("User").First(&article, id)
	return article
}
