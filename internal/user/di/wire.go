package di

import (
	"github.com/ahmadammarm/inventory-backend/internal/user/handler"
	"github.com/ahmadammarm/inventory-backend/internal/user/repos"
	"github.com/ahmadammarm/inventory-backend/internal/user/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedFakeUser(db *gorm.DB, val *validator.Validate) *handler.UserHandler {
    wire.Build(
        handler.NewUserHandler,
        service.NewUserService,
        repos.NewUserRepository,
    )

    return &handler.UserHandler{}
}