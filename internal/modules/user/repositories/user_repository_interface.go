package repositories

import UserModel "blog/internal/modules/user/models"

type UserRepositoryInterface interface {
	Register(user UserModel.User) (UserModel.User, error)
}
