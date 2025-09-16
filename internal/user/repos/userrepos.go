package repos

import (
	"github.com/ahmadammarm/inventory-backend/internal/user/model"
	"github.com/ahmadammarm/inventory-backend/pkg"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepos interface {
	SignupUser(user *model.User) error
}

type UserReposImpl struct {
	Database *gorm.DB
}

func (repo *UserReposImpl) SignupUser(user *model.User) error {

	hashPassword, err := pkg.HashPassword(user.Password)

	if err != nil {
		return err
	}

	newUser := model.User{
		ID:       uuid.New(),
		Name:     user.Name,
		Email:    user.Email,
		Password: hashPassword,
	}

	return repo.Database.Create(newUser).Error

}
