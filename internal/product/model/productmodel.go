package model

import (
	"time"

	"github.com/google/uuid"
)

type ProductCategory struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name" validate:"required,min=3"`
	Products  []Product `gorm:"foreignKey:ProductCategoryID" json:"products"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Product struct {
	ID                uuid.UUID       `gorm:"type:uuid;primaryKey" json:"id"`
	Name              string          `gorm:"type:varchar(100)" json:"name" validate:"required,min=3"`
	Price             float64         `json:"price" validate:"required,gte=0"`
	ProductCategoryID uuid.UUID       `gorm:"type:uuid;not null" json:"product_category_id" validate:"required"`
	ProductCategory   ProductCategory `gorm:"foreignKey:ProductCategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product_category"`
	CreatedAt         time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}
