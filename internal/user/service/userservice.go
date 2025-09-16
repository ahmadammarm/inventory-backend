package service

import (
	"errors"

	"github.com/ahmadammarm/inventory-backend/internal/user/model"
	userRepo "github.com/ahmadammarm/inventory-backend/internal/user/repos"
)

type UserService interface {
	SignupUser(user *model.User) error
}

type UserServiceImpl struct {
	userRepo userRepo.UserRepos
}

func (service *UserServiceImpl) SignupUser(user *model.User) error {

	if user.Email == "" || user.Password == "" {
		return errors.New("email and password must not be empty")
	}

	exists, err := service.userRepo.IsEmailExists(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}

	return service.userRepo.SignupUser(user)

}

func NewUserService(userRepo userRepo.UserRepos) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}
