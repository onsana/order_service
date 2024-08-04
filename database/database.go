package database

import (
	"fmt"
	"github.com/onsana/order_service/data/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBConnection() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=postgresdb user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	// Ensure the uuid-ossp extension is enabled
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		log.Fatal("Failed to create uuid-ossp extension. \n", err)
	}

	log.Println("running migrations")
	err = db.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Order{},
		&model.Address{})
	if err != nil {
		log.Fatal("Failed to auto migrate in database. \n", err)
	}

	return db
}
