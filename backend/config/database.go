package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// dsn := "host=localhost user=postgres password=postgres dbname=mini_eventify port=5433 sslmode=disable"
	dsn := "host=localhost user=postgres password=12345 dbname=mini_eventify port=5433 sslmode=disable"


	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database")
	}

	DB = database
	fmt.Println("✅ Database connected successfully")
}
