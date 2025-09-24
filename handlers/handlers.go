package handlers

import (
	"bioskop-api/database"
	"bioskop-api/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func TambahBioskop(c *gin.Context) {
	var bioskop models.Bioskop
	if err := c.ShouldBindJSON(&bioskop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if bioskop.Nama == "" || bioskop.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi kosong"})
		return
	}

	result := database.DB.Create(&bioskop)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menambahkan bioskop ke database"})
		return
	}

	c.JSON(http.StatusCreated, bioskop)
}

func AmbilSemuaBioskop(c *gin.Context) {
	var bioskops []models.Bioskop
	database.DB.Find(&bioskops)
	c.JSON(http.StatusOK, bioskops)
}

func AmbilBioskopByID(c *gin.Context) {
	id := c.Param("id")
	var bioskop models.Bioskop

	if err := database.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, bioskop)
}