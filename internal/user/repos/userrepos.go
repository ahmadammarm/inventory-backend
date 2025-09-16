package repos

import (
	"github.com/ahmadammarm/inventory-backend/internal/user/model"
	"github.com/ahmadammarm/inventory-backend/pkg"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepos interface {
	SignupUser(user *model.User) error
	IsEmailExists(email string) (bool, error)
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

func (repo *UserReposImpl) IsEmailExists(email string) (bool, error) {
	var user model.User
	err := repo.Database.Where("email = ?", email).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func NewUserRepository(database *gorm.DB) UserRepos {
	return &UserReposImpl{Database: database}
}
