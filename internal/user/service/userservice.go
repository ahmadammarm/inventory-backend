package service

import (
	"errors"

	"github.com/ahmadammarm/inventory-backend/internal/user/model"
	userRepo "github.com/ahmadammarm/inventory-backend/internal/user/repos"
	"github.com/ahmadammarm/inventory-backend/pkg/generatejwt"
	"github.com/ahmadammarm/inventory-backend/pkg/hashpassword"
)

type UserService interface {
	SignupUser(user *model.User) error
	SigninUser(userReq *model.User) (*model.User, string, error)
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

func (service *UserServiceImpl) SigninUser(userReq *model.User) (*model.User, string, error) {

    if userReq.Email == "" || userReq.Password == "" {
        return nil, "", errors.New("email and password are required")
    }


	user, err := service.userRepo.SigninUser(userReq.Email)

	if err != nil {
		return nil, "", err
	}


	if user == nil {
		return nil, "", errors.New("user not found")
	}


	if !hashpassword.IsPasswordMatch(userReq.Password, user.Password) {
		return nil, "", errors.New("wrong password")
	}
    

	token, err := generatejwt.GenerateJWT(user.Email)

	if err != nil {
		return nil, "", errors.New("failed to create the token")
	}

	return user, token, nil
}

func NewUserService(userRepo userRepo.UserRepos) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}
