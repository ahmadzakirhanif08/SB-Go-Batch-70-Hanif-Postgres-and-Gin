package main

import (
	"bioskop-api/database"
	"bioskop-api/handlers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	database.Connect()

	r := gin.Default()

	basicAuthUser := os.Getenv("BASIC_AUTH_USER")
	basicAuthPassword := os.Getenv("BASIC_AUTH_PASSWORD")

	r.GET("/bioskop", handlers.AmbilSemuaBioskop)
	r.GET("/bioskop/:id", handlers.AmbilBioskopByID)

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		basicAuthUser : basicAuthPassword,
	}))
	authorized.POST("/bioskop", handlers.TambahBioskop)
	authorized.PUT("/bioskop/:id", handlers.PerbaharuiBioskop)
	authorized.DELETE("/bioskop/:id", handlers.HapusBioskop)
	authorized.POST("/bioskop/:id/film", handlers.TambahFilm)

	//r.SetTrustedProxies([]string{"127.0.0.1"})
	//r.Run(":9090")

	port := os.Getenv("DB_PORT")
	if port == ""{
		port = "9090"
	}

	r.Run(":"+port)
}
