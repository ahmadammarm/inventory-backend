package repos

import (
	"github.com/ahmadammarm/inventory-backend/internal/product/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepos interface {
	CreateProduct(product *model.Product) error
}

type ProductReposImpl struct {
	Database *gorm.DB
}

func (repo *ProductReposImpl) CreateProduct(product *model.Product) error {
	product.ID = uuid.New()
	return repo.Database.Create(product).Error
}
