package migration

import (
	ArticleModel "blog/internal/modules/article/models"
	UserModel "blog/internal/modules/user/models"
	"blog/pkg/database"
	"fmt"
	"log"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&UserModel.User{}, ArticleModel.Article{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("migrate success")
}
