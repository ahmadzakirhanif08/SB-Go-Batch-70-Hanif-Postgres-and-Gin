package database

import (
	"bioskop-api/models"
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dbAddress := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSslMode := "disable"
	dbTimezone := "Asia/Jakarta"

	database := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbAddress, dbUser, dbPassword, dbName, dbPort, dbSslMode, dbTimezone,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(database), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error, please check connection")
	}

	fmt.Printf("Successfully connected to database: %s", dbName)

	DB.AutoMigrate(&models.Bioskop{}, &models.Film{}, &models.JadwalFilm{})
}
