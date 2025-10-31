package services

import (
	UserModel "blog/internal/modules/user/models"
	"blog/internal/modules/user/repositories"
	"blog/internal/modules/user/requests"
	userResponse "blog/internal/modules/user/responses"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repositories.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: repositories.New(),
	}

}

func (u *UserService) Register(request requests.RegisterRequest) (userResponse.User, error) {
	var response userResponse.User
	var user UserModel.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(3)
		return response, errors.New("error hashing password")
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Password = string(hashedPassword)

	result, err := u.userRepository.Register(user)
	if err != nil {
		log.Println(err.Error())
		return response, errors.New(err.Error())
	}
	return userResponse.ToUser(result), nil
}
