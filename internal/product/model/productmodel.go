package model

import (
	"time"

	"github.com/ahmadammarm/inventory-backend/internal/category/model"
	"github.com/google/uuid"
)

type Product struct {
	ID         uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Name       string         `gorm:"type:varchar(100)" json:"name" validate:"required,min=3"`
	ImageUrl   string         `gorm:"type:varchar(255)" json:"image_url" validate:"required"`
	Price      float64        `json:"price" validate:"required,gte=0"`
	CategoryID uuid.UUID      `gorm:"type:uuid;not null" json:"category_id" validate:"required"`
	Category   model.Category `gorm:"foreignKey:CategoryID" json:"category"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}
