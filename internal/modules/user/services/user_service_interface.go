package services

import (
	"blog/internal/modules/user/requests"
	userResponse "blog/internal/modules/user/responses"
)

type UserServiceInterface interface {
	Register(request requests.RegisterRequest) (userResponse.User, error)
}
