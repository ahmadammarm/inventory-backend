package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt int64     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64     `gorm:"autoUpdateTime" json:"updated_at"`
}