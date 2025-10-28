package responses

import (
	ArticleModel "blog/internal/modules/article/models"
	userResponse "blog/internal/modules/user/responses"
	"fmt"
)

type Article struct {
	ID        uint
	Title     string
	Body      string
	Image     string
	CreatedAt string
	User      userResponse.User
}

type Articles struct {
	Data []Article
}

func ToArticle(article ArticleModel.Article) Article {
	return Article{
		ID:        article.ID,
		Title:     article.Title,
		Body:      article.Body,
		Image:     fmt.Sprintf("/assets/img/demopic/10.jpg", article.Title),
		CreatedAt: fmt.Sprintf("%d/%02d/%02d", article.CreatedAt.Year(), article.CreatedAt.Month(), article.CreatedAt.Day()),
		User:      userResponse.ToUser(article.User),
	}
}

func ToArticles(articles []ArticleModel.Article) Articles {
	var response Articles
	for _, article := range articles {
		response.Data = append(response.Data, ToArticle(article))
	}
	return response
}
