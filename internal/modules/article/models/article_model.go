package models

import (
	"blog/internal/modules/user/models"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title  string `gorm:"varchar(190);not null"`
	Body   string `gorm:"text;not null"`
	UserID uint
	User   models.User
}
