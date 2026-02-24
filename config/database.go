package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=golang_clean port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed connect database")
	}

	DB = db
	fmt.Println("Database connected")
}
