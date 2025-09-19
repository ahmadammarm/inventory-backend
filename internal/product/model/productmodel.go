package model

import (
	"time"

	"github.com/google/uuid"
)

type ProductCategory struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Products  []Product `gorm:"foreignKey:ProductCategoryID" json:"products"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}


type Product struct {
	ID                uuid.UUID       `gorm:"type:uuid;primaryKey" json:"id"`
	Name              string          `gorm:"type:varchar(100)" json:"name"`
	Price             float64         `json:"price"`
	ProductCategoryID uuid.UUID       `gorm:"type:uuid;not null" json:"product_category_id"`
	ProductCategory   ProductCategory `gorm:"foreignKey:ProductCategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"product_category"`
	CreatedAt         time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
}
