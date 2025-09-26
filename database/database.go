package database

import (
	"fmt"
	"log"
	"bioskop-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const(
	dbAddress = "localhost"
	dbUser = "admin"
	dbPassword = "admin"
	dbName = "golang"
	dbPort = "5432"
	dbSslMode = "disable"
	dbTimezone = "Asia/Jakarta"
)

var DB *gorm.DB

func Connect() {
	database := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbAddress, dbUser, dbPassword, dbName, dbPort, dbSslMode, dbTimezone,
	)
	var err error
	DB, err = gorm.Open(postgres.Open(database), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error, please check connection")
	}

	fmt.Printf("Successfully connected to database: %s",dbName)
	
	DB.AutoMigrate(&models.Bioskop{}, &models.Film{}, &models.JadwalFilm{})
}