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
	var database string
	if os.Getenv("DATABASE_URL") != "" {
		// Jika ada, gunakan DSN dari Railway secara langsung
		database = os.Getenv("DATABASE_URL")
	} else {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")
		
		database = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			dbHost, dbUser, dbPassword, dbName, dbPort)
	}
	var err error
	DB, err = gorm.Open(postgres.Open(database), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection error, please check connection")
	}

	fmt.Println("Successfully connected to database")

	DB.AutoMigrate(&models.Bioskop{}, &models.Film{}, &models.JadwalFilm{})
}
