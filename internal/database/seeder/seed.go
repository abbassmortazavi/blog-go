package seeder

import (
	ArticleModel "blog/internal/modules/article/models"
	UserModel "blog/internal/modules/user/models"
	"blog/pkg/database"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := database.Connection()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	user := UserModel.User{
		Name:     "Bob",
		Email:    "abbassmortazavi@gmail.com",
		Password: string(hashedPassword),
	}
	db.Create(&user)
	log.Printf("User %v has been created", user.Name)

	for i := 0; i < 10; i++ {
		article := ArticleModel.Article{
			Title:  fmt.Sprintf("Title %d", i),
			Body:   fmt.Sprintf("Body %d", i),
			UserID: user.ID,
		}
		db.Create(&article)
		log.Printf("Article %v has been created", article.Title)

	}
	log.Println("All articles have been seeded")
}
