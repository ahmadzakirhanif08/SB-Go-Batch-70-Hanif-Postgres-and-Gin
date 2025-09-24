package main

import (
	"bioskop-api/database"
	"bioskop-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.GET("/bioskop", handlers.AmbilSemuaBioskop)
	r.GET("/bioskop/:id", handlers.AmbilBioskopByID)

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"gaktau": "bilangwow",
	}))
	authorized.POST("/bioskop", handlers.TambahBioskop)

	r.Run(":8080")
}