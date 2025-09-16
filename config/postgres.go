package config

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"golang.org/x/exp/slog"

	"github.com/ahmadammarm/inventory-backend/internal/user/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var (
	gormDB   *gorm.DB
	err      error
	syncOnce sync.Once
)

func PostgresConnect() (*gorm.DB, error) {
	syncOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			slog.Error("Error loading .env file", slog.String("error", err.Error()))
		}

		host := os.Getenv("DB_HOST")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")
		port := os.Getenv("DB_PORT")


		dsnNoDB := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, port)
		sqlDB, err := sql.Open("postgres", dsnNoDB)
		if err != nil {
			slog.Error("Failed to connect to Postgres server", slog.String("error", err.Error()))
			return
		}
		defer sqlDB.Close()


		var exists bool
		checkQuery := fmt.Sprintf("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = '%s')", dbname)
		if err = sqlDB.QueryRow(checkQuery).Scan(&exists); err != nil {
			slog.Error("Failed to check database existence", slog.String("error", err.Error()))
			return
		}

		if !exists {
			createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbname)
			if _, err = sqlDB.Exec(createDBQuery); err != nil {
				slog.Error("Failed to create database", slog.String("error", err.Error()))
				return
			}
			slog.Info("Database created successfully", slog.String("dbname", dbname))
		}


		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			host, user, password, dbname, port)

		gormDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			slog.Error("Failed to connect to database", slog.String("error", err.Error()))
			return
		}

		dbPool, err := gormDB.DB()
		if err != nil {
			slog.Error("Failed to get DB pool", slog.String("error", err.Error()))
			return
		}
		dbPool.SetMaxIdleConns(10)
		dbPool.SetMaxOpenConns(100)

		if err = gormDB.AutoMigrate(&model.User{}); err != nil {
			slog.Error("Auto migration failed", slog.String("error", err.Error()))
			return
		}

        
		slog.Info("Auto migration completed")
		slog.Info("Database connection established")
	})

	return gormDB, err
}
