package database

import (
	"fmt"
	"log"
	"bioskop-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database := "host=localhost user=admin password=admin dbname=golang port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error
	DB, err = gorm.Open(postgres.Open(database), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database")
	}

	fmt.Println("Berhasil terhubung ke database!")
	
	DB.AutoMigrate(&models.Bioskop{})
}