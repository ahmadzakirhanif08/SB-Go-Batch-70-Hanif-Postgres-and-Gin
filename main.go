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
	authorized.PUT("/bioskop/:id", handlers.PerbaharuiBioskop)
	authorized.DELETE("/bioskop/:id", handlers.HapusBioskop)

	r.Run(":9090")
}
