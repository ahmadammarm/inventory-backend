package repos

import (
	"github.com/ahmadammarm/inventory-backend/internal/category/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryRepos interface {
	CreateCategory(category *model.Category) error
}

type CategoryReposImpl struct {
	Database *gorm.DB
}

func (repo *CategoryReposImpl) CreateCategory(category *model.Category) error {
	category.ID = uuid.New()
	return repo.Database.Create(category).Error
}
