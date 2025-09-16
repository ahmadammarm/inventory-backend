package main

import (
	"github.com/ahmadammarm/inventory-backend/config"
	"golang.org/x/exp/slog"
)

func main() {
    database, err := config.PostgresConnect()

    if err != nil {
        slog.Error("Failed to connected with database")
    }

    database.Logger.LogMode(1)
}