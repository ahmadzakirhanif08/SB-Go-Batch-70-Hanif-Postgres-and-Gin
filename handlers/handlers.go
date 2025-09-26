package handlers

import (
	"bioskop-api/database"
	"bioskop-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TambahBioskop(c *gin.Context) {
	var bioskop models.Bioskop
	if err := c.ShouldBindJSON(&bioskop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if bioskop.Nama == "" || bioskop.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "field name and location empty"})
		return
	}

	result := database.DB.Create(&bioskop)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed add bioskop to database"})
		return
	}

	c.JSON(http.StatusCreated, bioskop)
}

func AmbilSemuaBioskop(c *gin.Context) {
	var bioskops []models.Bioskop

	if err := database.DB.Preload("FilmTayang").Preload("FilmTayang.Jadwal").Find(&bioskops).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed fetching data from film bioskop."})
		return
	}

	c.JSON(http.StatusOK, bioskops)
}

func AmbilBioskopByID(c *gin.Context) {
	id := c.Param("id")
	var bioskop models.Bioskop

	if err := database.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop not found"})
		return
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not valid. ID should be integer."})
		return
	}

	if err := database.DB.Preload("FilmTayang").Preload("FilmTayang.Jadwal").First(&bioskop, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop not found."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed fetching data, server error."})
		}
		return
	}

	if err := database.DB.First(&bioskop, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch bioskop data."})
		}
		return
	}

	c.JSON(http.StatusOK, bioskop)
}

func PerbaharuiBioskop(c *gin.Context) {
	idStr := c.Param("id")
	var bioskop models.Bioskop

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Not Valid. ID must integer."})
		return
	}

	if err := database.DB.First(&bioskop, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop not found."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed fetching data."})
		}
		return
	}

	if err := c.ShouldBindJSON(&bioskop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request not valid, check JSON."})
		return
	}

	if bioskop.Nama == "" || bioskop.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama and Lokasi empty."})
		return
	}

	if err := database.DB.Save(&bioskop).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed update data."})
		return
	}

	c.JSON(http.StatusOK, bioskop)
}

func HapusBioskop(c *gin.Context) {
	id := c.Param("id")
	var bioskop models.Bioskop

	_, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Not Valid. ID must integer."})
		return
	}

	if err := database.DB.First(&bioskop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "bioskop not found"})
		return
	}

	if err := database.DB.First(&bioskop, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "bioskop not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed fetching data."})
		}
		return
	}

	database.DB.Delete(&bioskop)
	c.JSON(http.StatusOK, gin.H{"message": "delete success"})
}

func TambahFilm(c *gin.Context) {
	idStr := c.Param("id")

	var film models.Film
	var bioskop models.Bioskop

	bioskopID, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID bioskop is not valid"})
		return
	}

	if err := database.DB.First(&bioskop, bioskopID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Bioskop not found ditemukan."})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed fetching data."})
		}
		return
	}

	if err := c.ShouldBindJSON(&film); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "request not valid"})
		return
	}

	if film.NamaFilm == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama Film can't be empty"})
		return
	}

	film.BioskopID = uint(bioskopID)

	if err := database.DB.Create(&film).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed add film, please try again"})
	}

	c.JSON(http.StatusCreated, film)
}
