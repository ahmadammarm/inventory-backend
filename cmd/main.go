package main

import (
	"log"
	"os"

	"github.com/ahmadammarm/inventory-backend/config"
	users "github.com/ahmadammarm/inventory-backend/internal/user/di"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

func init() {
    config.SupabaseStorageConnect()
}

func main() {
    database, err := config.PostgresConnect()

    if err != nil {
        slog.Error("Failed to connected with database")
    }

    app := fiber.New(fiber.Config{
        DisableStartupMessage: true,
    })

    api := app.Group("/api")
    users.InitializedUser(database, validator.New()).UserRouters(api)

    database.Logger.LogMode(1)

    if error := app.Listen(":8080"); error != nil {
		log.Printf("Failed to start server: %v", error)
		os.Exit(1)
	}
}