package database

import (
	entity "chat-service/internal/infrastructure/database/entity"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	databaseName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")

	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, databaseName, port)

	db, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(entity.Message{})

	return db, nil
}