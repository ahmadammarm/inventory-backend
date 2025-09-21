package repos

import (
	"errors"

	"github.com/ahmadammarm/inventory-backend/internal/product/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepos interface {
	CreateCategory(product *model.ProductCategory) error
    IsProductCategoryExists(name string) (bool, error)
	GetAllCategories() ([]model.ProductCategory, error)
	GetCategoryByID(id string) (*model.ProductCategory, error)
	UpdateCategory(product *model.ProductCategory) error
	DeleteCategory(id string) error
}

type ProductReposImpl struct {
	Database *gorm.DB
}

func (repo *ProductReposImpl) CreateCategory(category *model.ProductCategory) error {

	category.ID = uuid.New()
	return repo.Database.Create(category).Error

}


func (repo *ProductReposImpl) IsProductCategoryExists(name string) (bool, error) {
    var category model.ProductCategory
    err := repo.Database.Where("name = ?", name).First(&category).Error
    if err == nil {
        return true, nil
    }
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return false, nil
    }
    return false, err
}

