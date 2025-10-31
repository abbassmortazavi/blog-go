package repositories

import (
	UserModel "blog/internal/modules/user/models"
	"blog/pkg/database"
	"errors"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository {
	return &UserRepository{
		DB: database.Connection(),
	}
}

func (r *UserRepository) Register(user UserModel.User) (UserModel.User, error) {
	var newUser UserModel.User
	if err := r.DB.Where("email = ?", user.Email).First(&user).Error; err == nil {
		log.Println(1)
		return newUser, errors.New("user already exists")
	}
	if err := r.DB.Create(&user).Scan(&newUser).Error; err != nil {
		log.Println(2)
		return newUser, err
	}
	log.Println("this repository=>>>>>>>", newUser)
	return newUser, nil
}
